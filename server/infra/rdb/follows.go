//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package rdb

import (
	"context"
	"tMinamiii/Tweet/domain"

	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type Follows interface {
	CreateTx(ctx context.Context, tx dbr.SessionRunner, userID, followUserID int64) (int64, error)
	DeleteTx(ctx context.Context, tx dbr.SessionRunner, userID, followUserID int64) error
	LoadByUserID(ctx context.Context, userID int64) (*[]domain.Follow, error)
	LoadByUserIDAndFollowUserIDs(ctx context.Context, userID int64, followUserIDs []int64) (*[]domain.Follow, error)
	LoadByUserIDAndFollowUserID(ctx context.Context, userID, followUserID int64) (*domain.Follow, error)
	LoadByUserIDAndFollowUserIDTx(ctx context.Context, tx dbr.SessionRunner, userID, followUserID int64) (*domain.Follow, error)
}

type followsTable struct{}

func NewFollowsTable() Follows {
	return followsTable{}
}

func (f followsTable) CreateTx(ctx context.Context, tx dbr.SessionRunner, userID, followUserID int64) (int64, error) {
	result, err := tx.InsertInto("follows").
		Pair("user_id", userID).
		Pair("follow_user_id", followUserID).
		ExecContext(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create follows record")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get last inserted id")
	}

	return id, nil
}

func (f followsTable) DeleteTx(ctx context.Context, tx dbr.SessionRunner, userID, followUserID int64) error {
	_, err := tx.DeleteFrom("follows").
		Where("user_id = ? ", userID).
		Where("follow_user_id = ? ", followUserID).
		ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to delete follows record")
	}

	return nil
}

// LoadByUserID フォローしているユーザーを探すための関数
func (f followsTable) LoadByUserID(ctx context.Context, userID int64) (*[]domain.Follow, error) {
	sess := GetTweetSession()
	m := &[]domain.Follow{}
	_, err := sess.
		Select("*").
		From("follows").
		Where("user_id = ?", userID).
		LoadContext(ctx, m)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load follow records by user id")
	}

	return m, nil
}

// LoadByUserIDAndFollowUserIDs フォロー済みのユーザー一覧を取得する関数
func (f followsTable) LoadByUserIDAndFollowUserIDs(ctx context.Context, userID int64, followUserIDs []int64) (*[]domain.Follow, error) {
	sess := GetTweetSession()
	m := &[]domain.Follow{}
	_, err := sess.
		Select("*").
		From("follows").
		Where("user_id = ? AND follow_user_id IN ?", userID, followUserIDs).
		LoadContext(ctx, m)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load follow records by user id and follow user ids")
	}

	return m, nil
}

// LoadByUserIDAndFollowUserID レコードを1件にしぼって取得する関数
func (f followsTable) LoadByUserIDAndFollowUserID(ctx context.Context, userID, followUserID int64) (*domain.Follow, error) {
	sess := GetTweetSession()
	return f.LoadByUserIDAndFollowUserIDTx(ctx, sess, userID, followUserID)
}

// LoadByUserIDAndFollowUserIDTx レコードを1件にしぼって取得する関数
func (f followsTable) LoadByUserIDAndFollowUserIDTx(ctx context.Context, tx dbr.SessionRunner, userID, followUserID int64) (*domain.Follow, error) {
	m := &domain.Follow{}
	_, err := tx.
		Select("*").
		From("follows").
		Where("user_id = ? AND follow_user_id = ?", userID, followUserID).
		LoadContext(ctx, m)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return m, nil
		}
		return nil, errors.Wrap(err, "failed to load follow records by user id and follow user id")
	}

	return m, nil
}
