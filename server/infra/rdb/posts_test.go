package rdb

import (
	"context"
	"tMinamiii/Tweet/domain"
	"testing"
	"time"

	"github.com/gocraft/dbr"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_postsTable_CreateTx(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE posts") })

	type args struct {
		ctx     context.Context
		tx      dbr.SessionRunner
		userID  int64
		content string
	}
	tests := []struct {
		name        string
		args        args
		wantUserID  int64
		wantContent string
		wantErr     bool
		timeout     time.Duration
	}{
		{
			name:        "postレコード作成成功",
			args:        args{ctx: context.Background(), tx: GetTweetSession(), userID: 1, content: "tweet tweet"},
			wantUserID:  1,
			wantContent: "tweet tweet",
			timeout:     30 * time.Second,
		},
		{
			name:    "タイムアウトエラー",
			args:    args{ctx: context.Background(), tx: GetTweetSession(), userID: 1, content: "tweet tweet"},
			timeout: -time.Second,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			p := &postsTable{}
			got, err := p.CreateTx(ctx, tt.args.tx, tt.args.userID, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("postsTable.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			wantRecord := &domain.Post{UUID: got, UserID: tt.wantUserID, Content: tt.wantContent}
			gotRecord := &domain.Post{}
			sess.Select("BIN_TO_UUID(uuid) AS uuid, user_id, content").
				From("posts").
				Where("uuid = UUID_TO_BIN(?)", got).
				LoadOneContext(ctx, gotRecord)
			ignore := cmpopts.IgnoreFields(domain.Post{}, "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(gotRecord, wantRecord, ignore); diff != "" {
				t.Errorf("postsTable.Create() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_postsTable_LoadByUUIDTx(t *testing.T) {
	sess := GetTweetSession()
	t.Cleanup(func() { sess.Exec("TRUNCATE TABLE posts") })

	uuid := "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f"
	userID := int64(1)
	content := "tweet"
	ctx := context.Background()
	sess.
		InsertBySql("INSERT INTO posts (uuid, user_id, content) VALUES (UUID_TO_BIN(?), ?, ?)", uuid, userID, content).
		ExecContext(ctx)

	type args struct {
		ctx  context.Context
		tx   dbr.SessionRunner
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Post
		wantErr bool
		timeout time.Duration
	}{
		{
			name: "レコードを取得",
			args: args{
				ctx:  context.Background(),
				tx:   GetTweetSession(),
				uuid: "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f",
			},
			want: &domain.Post{
				UUID:    uuid,
				UserID:  userID,
				Content: content,
			},
			timeout: 30 * time.Second,
		},
		{
			name: "レコードが存在しないため空構造体",
			args: args{
				ctx:  context.Background(),
				tx:   GetTweetSession(),
				uuid: "0189f7ea-ae2c-7809-8aeb-000000000000",
			},
			want:    &domain.Post{},
			timeout: 30 * time.Second,
		},
		{
			name: "タイムアウトエラー",
			args: args{
				ctx:  context.Background(),
				tx:   GetTweetSession(),
				uuid: "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f",
			},
			wantErr: true,
			timeout: -time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(tt.args.ctx, tt.timeout)
			defer cancel()

			p := &postsTable{}
			got, err := p.LoadByUUIDTx(ctx, tt.args.tx, tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("postsTable.LoadByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ignore := cmpopts.IgnoreFields(domain.Post{}, "CreatedAt", "UpdatedAt")
			if diff := cmp.Diff(got, tt.want, ignore); diff != "" {
				t.Errorf("postsTable.LoadByUUID() got diff: (-got +want)\n%s", diff)
			}
		})
	}
}
