//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package rdb

import (
	"context"
	"tMinamiii/Tweet/domain"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_postsUsersTable_LoadByUserIDs(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() {
		sess.Exec("TRUNCATE TABLE posts")
		sess.Exec("TRUNCATE TABLE users")
	})

	ctx := context.Background()
	p := &postsTable{}
	uuid1, _ := p.CreateTx(ctx, sess, 1, "tweet 1")
	uuid2, _ := p.CreateTx(ctx, sess, 2, "tweet 2")
	uuid3, _ := p.CreateTx(ctx, sess, 3, "tweet 3")
	uuid4, _ := p.CreateTx(ctx, sess, 1, "tweet 4")
	uuid5, _ := p.CreateTx(ctx, sess, 2, "tweet 5")
	uuid6, _ := p.CreateTx(ctx, sess, 3, "tweet 6")

	accountID1 := "test_account_id_1"
	accountID2 := "test_account_id_2"
	accountID3 := "test_account_id_3"
	username1 := "test_username_1"
	username2 := "test_username_2"
	username3 := "test_username_3"
	sess.InsertInto("users").Pair("id", 1).Pair("account_id", accountID1).Pair("username", username1).ExecContext(ctx)
	sess.InsertInto("users").Pair("id", 2).Pair("account_id", accountID2).Pair("username", username2).ExecContext(ctx)
	sess.InsertInto("users").Pair("id", 3).Pair("account_id", accountID3).Pair("username", username3).ExecContext(ctx)

	type args struct {
		ctx       context.Context
		userIDs   []int64
		limit     int
		sinceUUID *string
	}
	tests := []struct {
		name    string
		args    args
		want    *[]domain.PostUser
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "全レコードを取得",
			args: args{
				ctx:       context.Background(),
				userIDs:   []int64{1, 2, 3},
				limit:     6,
				sinceUUID: nil,
			},
			want: &[]domain.PostUser{
				{UUID: uuid6, UserID: 3, Username: username3, AccountID: accountID3, Content: "tweet 6"},
				{UUID: uuid5, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 5"},
				{UUID: uuid4, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 4"},
				{UUID: uuid3, UserID: 3, Username: username3, AccountID: accountID3, Content: "tweet 3"},
				{UUID: uuid2, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 2"},
				{UUID: uuid1, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 1"},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "全レコードを取得(sinceUUIDが空文字)",
			args: args{
				ctx:       context.Background(),
				userIDs:   []int64{1, 2, 3},
				limit:     6,
				sinceUUID: nil,
			},
			want: &[]domain.PostUser{
				{UUID: uuid6, UserID: 3, Username: username3, AccountID: accountID3, Content: "tweet 6"},
				{UUID: uuid5, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 5"},
				{UUID: uuid4, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 4"},
				{UUID: uuid3, UserID: 3, Username: username3, AccountID: accountID3, Content: "tweet 3"},
				{UUID: uuid2, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 2"},
				{UUID: uuid1, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 1"},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "レコードを3件取得",
			args: args{
				ctx:       context.Background(),
				userIDs:   []int64{1, 2, 3},
				limit:     3,
				sinceUUID: nil,
			},
			want: &[]domain.PostUser{
				{UUID: uuid6, UserID: 3, Username: username3, AccountID: accountID3, Content: "tweet 6"},
				{UUID: uuid5, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 5"},
				{UUID: uuid4, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 4"},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "user_id 1, 2の投稿を取得",
			args: args{
				ctx:       context.Background(),
				userIDs:   []int64{1, 2},
				limit:     4,
				sinceUUID: nil,
			},

			want: &[]domain.PostUser{
				{UUID: uuid5, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 5"},
				{UUID: uuid4, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 4"},
				{UUID: uuid2, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 2"},
				{UUID: uuid1, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 1"},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "ページネーション uuid3移行のレコードを取得",
			args: args{
				ctx:       context.Background(),
				userIDs:   []int64{1, 2, 3},
				limit:     10,
				sinceUUID: &uuid3,
			},
			want: &[]domain.PostUser{
				{UUID: uuid2, UserID: 2, Username: username2, AccountID: accountID2, Content: "tweet 2"},
				{UUID: uuid1, UserID: 1, Username: username1, AccountID: accountID1, Content: "tweet 1"},
			},
			timeout: 30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:       context.Background(),
				userIDs:   []int64{1, 2, 3},
				limit:     6,
				sinceUUID: nil,
			},
			wantErr: true,
			timeout: -1 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			p := &postsUsersTable{}
			got, err := p.LoadByUserIDs(ctx, tt.args.userIDs, tt.args.limit, tt.args.sinceUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("postsUsersTable.LoadByUserIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignore := cmpopts.IgnoreFields(domain.PostUser{}, "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("postsUsersTable.LoadByUserIDs() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}
