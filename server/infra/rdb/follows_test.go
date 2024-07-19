package rdb

import (
	"context"
	"tMinamiii/Tweet/domain"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFollowsTable_Create(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE follows") })

	type args struct {
		ctx          context.Context
		userID       int64
		followUserID int64
	}
	tests := []struct {
		name             string
		args             args
		wantUserID       int64
		wantFollowUserID int64
		wantErr          bool
		timeout          time.Duration
	}{
		{
			name: "followレコード作成に成功",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			wantUserID:       1,
			wantFollowUserID: 10,
			timeout:          30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			sess := GetTweetSession()
			tx, _ := sess.Begin()
			defer tx.RollbackUnlessCommitted()

			f := &followsTable{}
			got, err := f.CreateTx(ctx, tx, tt.args.userID, tt.args.followUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowsTable.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantRecord := &domain.Follow{ID: got, UserID: tt.wantUserID, FollowUserID: tt.wantFollowUserID}
			gotRecord := &domain.Follow{}
			tx.Select("*").
				From("follows").
				Where("id = ?", got).
				LoadOneContext(ctx, gotRecord)
			ignore := cmpopts.IgnoreFields(domain.Follow{}, "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(gotRecord, wantRecord, ignore); diff != "" {
				t.Errorf("followsTable.Create() got diff: (-got +want)\n%s", diff)
			}
			tx.Rollback()
		})
	}
}

func TestFollowsTable_LoadByUserID(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE follows") })

	f := &followsTable{}
	f.CreateTx(context.Background(), sess, 1, 10)
	f.CreateTx(context.Background(), sess, 1, 20)
	f.CreateTx(context.Background(), sess, 2, 10)

	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    *[]domain.Follow
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "user_id 1 がフォローしているユーザー取得",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: &[]domain.Follow{
				{UserID: 1, FollowUserID: 10},
				{UserID: 1, FollowUserID: 20},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			f := &followsTable{}
			got, err := f.LoadByUserID(ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowsTable.LoadByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignore := cmpopts.IgnoreFields(domain.Follow{}, "ID", "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("followsTable.LoadByUserID()() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFollowsTable_LoadByUserIDAndFollowUserIDs(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE follows") })

	f := &followsTable{}
	f.CreateTx(context.Background(), sess, 1, 10)
	f.CreateTx(context.Background(), sess, 1, 20)
	f.CreateTx(context.Background(), sess, 2, 20)

	type args struct {
		ctx           context.Context
		userID        int64
		followUserIDs []int64
	}
	tests := []struct {
		name    string
		f       *followsTable
		args    args
		want    *[]domain.Follow
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "user id 1 がフォローしているレコード一覧を取得",
			args: args{
				ctx:           context.Background(),
				userID:        1,
				followUserIDs: []int64{10, 20, 30, 40},
			},
			want: &[]domain.Follow{
				{UserID: 1, FollowUserID: 10},
				{UserID: 1, FollowUserID: 20},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:           context.Background(),
				userID:        1,
				followUserIDs: []int64{10, 20, 30, 40},
			},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			f := &followsTable{}
			got, err := f.LoadByUserIDAndFollowUserIDs(ctx, tt.args.userID, tt.args.followUserIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowsTable.LoadByUserIDAndFollowUserIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignore := cmpopts.IgnoreFields(domain.Follow{}, "ID", "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("FollowsTable.LoadByUserIDAndFollowUserIDs() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFollowsTable_LoadByUserIDAndFollowUserID(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE follows") })

	f := &followsTable{}
	f.CreateTx(context.Background(), sess, 1, 10)

	type args struct {
		ctx          context.Context
		userID       int64
		followUserID int64
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Follow
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "フォロー済みのレコードを取得",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			want:    &domain.Follow{UserID: 1, FollowUserID: 10},
			timeout: 30 * time.Second,
		},
		{
			name: "レコードが0件の場合",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 100,
			},
			want:    &domain.Follow{},
			timeout: 30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			f := &followsTable{}
			got, err := f.LoadByUserIDAndFollowUserID(ctx, tt.args.userID, tt.args.followUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowsTable.LoadByUserIDAndFollowUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignore := cmpopts.IgnoreFields(domain.Follow{}, "ID", "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("FollowsTable.LoadByUserIDAndFollowUserID() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFollowsTable_LoadByUserIDAndFollowUserIDTx(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE follows") })

	f := &followsTable{}
	f.CreateTx(context.Background(), sess, 1, 10)

	type args struct {
		ctx          context.Context
		userID       int64
		followUserID int64
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Follow
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "フォロー済みのレコードを取得",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			want:    &domain.Follow{UserID: 1, FollowUserID: 10},
			timeout: 30 * time.Second,
		},
		{
			name: "レコードが0件の場合",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 100,
			},
			want:    &domain.Follow{},
			timeout: 30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			f := &followsTable{}
			sess := GetTweetSession()
			got, err := f.LoadByUserIDAndFollowUserIDTx(ctx, sess, tt.args.userID, tt.args.followUserID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowsTable.LoadByUserIDAndFollowUserIDTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignore := cmpopts.IgnoreFields(domain.Follow{}, "ID", "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("FollowsTable.LoadByUserIDAndFollowUserIDTx() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_followsTable_DeleteTx(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE follows") })

	f := &followsTable{}
	// The line `// f.CreateTx(context.Background(), sess, 1, 10)` is a commented-out function call to the
	// `CreateTx` method of the `followsTable` struct. This method is used to create a new follow
	// relationship between a user with ID 1 and another user with ID 10 within a transaction.
	f.CreateTx(context.Background(), sess, 1, 10)

	type args struct {
		ctx          context.Context
		userID       int64
		followUserID int64
	}
	tests := []struct {
		name    string
		f       *followsTable
		args    args
		want    *domain.Follow
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "フォローレコード削除に成功",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			want:    &domain.Follow{},
			timeout: 30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:          context.Background(),
				userID:       1,
				followUserID: 10,
			},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			f := &followsTable{}
			tx, _ := sess.Begin()
			defer tx.RollbackUnlessCommitted()

			if err := f.DeleteTx(ctx, tx, tt.args.userID, tt.args.followUserID); (err != nil) != tt.wantErr {
				t.Errorf("followsTable.DeleteTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, _ := f.LoadByUserIDAndFollowUserIDTx(ctx, tx, tt.args.userID, tt.args.followUserID)

			ignore := cmpopts.IgnoreFields(domain.Follow{}, "ID", "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("FollowsTable.DeleteTx() got diff: (-got +want)\n%s", diff)
			}
			tx.Rollback()
		})
	}
}
