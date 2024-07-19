package rdb

import (
	"context"
	"tMinamiii/Tweet/domain"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_usersTable_LoadByAccountID(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE users") })
	sess.InsertInto("users").Pair("account_id", "test_account_id").Pair("username", "test_username").Exec()

	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "レコード取得に成功",
			args: args{ctx: context.Background(), accountID: "test_account_id"},
			want: &domain.User{
				ID:        1,
				AccountID: "test_account_id",
				Username:  "test_username",
			},
			timeout: 30 * time.Second,
		},
		{
			name:    "該当するレコードが0件",
			args:    args{ctx: context.Background(), accountID: "invalid_account_id"},
			want:    &domain.User{},
			timeout: 30 * time.Second,
		},
		{
			name:    "タイムアウトエラー",
			args:    args{ctx: context.Background(), accountID: "invalid_account_id"},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			u := &usersTable{}
			got, err := u.LoadByAccountID(ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersTable.LoadByAccountID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ignore := cmpopts.IgnoreFields(domain.User{}, "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("usersTable.LoadByAccountID() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_usersTable_FindByUsername(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE users") })
	sess.InsertInto("users").Pair("account_id", "test_account_id_01").Pair("username", "test_username_01").Exec()
	sess.InsertInto("users").Pair("account_id", "test_account_id_02").Pair("username", "test_username_02").Exec()
	sess.InsertInto("users").Pair("account_id", "test_account_id_03").Pair("username", "test_username_03").Exec()
	sess.InsertInto("users").Pair("account_id", "test_account_id_04").Pair("username", "test_name_04").Exec()
	sess.InsertInto("users").Pair("account_id", "test_account_id_05").Pair("username", "test_name_05").Exec()

	type args struct {
		ctx           context.Context
		exceptUserIDs []int64
		username      string
	}
	tests := []struct {
		name    string
		u       *usersTable
		args    args
		want    *[]domain.User
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "絞り込みに成功",
			args: args{ctx: context.Background(), exceptUserIDs: []int64{3}, username: "username"},
			want: &[]domain.User{
				{ID: 1, AccountID: "test_account_id_01", Username: "test_username_01"},
				{ID: 2, AccountID: "test_account_id_02", Username: "test_username_02"},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "除外ユーザー以外、全レコード取得に成功",
			args: args{ctx: context.Background(), exceptUserIDs: []int64{3}, username: ""},
			want: &[]domain.User{
				{ID: 1, AccountID: "test_account_id_01", Username: "test_username_01"},
				{ID: 2, AccountID: "test_account_id_02", Username: "test_username_02"},
				{ID: 4, AccountID: "test_account_id_04", Username: "test_name_04"},
				{ID: 5, AccountID: "test_account_id_05", Username: "test_name_05"},
			},
			timeout: 30 * time.Second,
		},
		{
			name:    "該当するレコードが0件",
			args:    args{ctx: context.Background(), exceptUserIDs: []int64{3}, username: "invalid_username"},
			want:    &[]domain.User{},
			timeout: 30 * time.Second,
		},
		{
			name:    "タイムアウトエラー",
			args:    args{ctx: context.Background(), exceptUserIDs: []int64{3}, username: "username"},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			u := &usersTable{}
			got, err := u.FindByUsername(ctx, tt.args.exceptUserIDs, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersTable.FindByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ignore := cmpopts.IgnoreFields(domain.User{}, "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("usersTable.FindByUsername() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_usersTable_LoadByID(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE users") })
	sess.InsertInto("users").Pair("id", 1).Pair("account_id", "test_account_id").Pair("username", "test_username").Exec()

	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "レコード取得に成功",
			args: args{ctx: context.Background(), userID: 1},
			want: &domain.User{
				ID:        1,
				AccountID: "test_account_id",
				Username:  "test_username",
			},
			timeout: 30 * time.Second,
		},
		{
			name:    "該当するレコードが0件",
			args:    args{ctx: context.Background(), userID: 100},
			want:    &domain.User{},
			timeout: 30 * time.Second,
		},
		{
			name:    "タイムアウトエラー",
			args:    args{ctx: context.Background(), userID: 1},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			u := &usersTable{}
			got, err := u.LoadByID(ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("usersTable.LoadByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignore := cmpopts.IgnoreFields(domain.User{}, "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("usersTable.LoadByID() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}
