//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package rdb

import (
	"context"
	"strings"
	"tMinamiii/Tweet/domain"

	"github.com/pkg/errors"
)

type PostsUsers interface {
	LoadByUserIDs(ctx context.Context, userIDs []int64, limit int, sinceUUID *string) (*[]domain.PostUser, error)
}

type postsUsersTable struct {
}

func NewPostsUsers() PostsUsers {
	return postsUsersTable{}
}

// LoadByUserIDs タイムラインを取得するための関数
// NOTE: 投稿数が増える、もしくはフォローユーザーが増えると遅くなる
func (p postsUsersTable) LoadByUserIDs(ctx context.Context, userIDs []int64, limit int, sinceUUID *string) (*[]domain.PostUser, error) {
	m := &[]domain.PostUser{}
	sess := GetTweetSession()
	stmt := sess.Select("BIN_TO_UUID(posts.uuid) AS uuid, posts.user_id, users.username, users.account_id, posts.content, posts.created_at, posts.updated_at").
		From("posts").
		Join("users", "posts.user_id = users.id").
		OrderDesc("uuid").
		Where("user_id IN ? ", userIDs)

	if sinceUUID != nil {
		if trimsSinceUUID := strings.TrimSpace(*sinceUUID); trimsSinceUUID != "" {
			stmt = stmt.Where("uuid < UUID_TO_BIN(?)", sinceUUID)
		}
	}

	_, err := stmt.
		Limit(uint64(limit)).
		OrderAsc("uuid").
		LoadContext(ctx, m)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load timeline records by user id")
	}

	return m, nil
}
