//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package rdb

import (
	"context"
	"tMinamiii/Tweet/domain"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type Posts interface {
	CreateTx(ctx context.Context, tx dbr.SessionRunner, userID int64, content string) (string, error)
	LoadByUUIDTx(ctx context.Context, tx dbr.SessionRunner, uuid string) (*domain.Post, error)
}

type postsTable struct {
}

func NewPostsTable() Posts {
	return postsTable{}
}

// Create つぶやきを保存する関数
func (p postsTable) CreateTx(ctx context.Context, tx dbr.SessionRunner, userID int64, content string) (string, error) {
	// 主キーには時系列ソートが可能なuuidv7を使用する
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate post record uuid")
	}

	_, err = tx.
		InsertBySql("INSERT INTO posts (uuid, user_id, content) VALUES (UUID_TO_BIN(?), ?, ?)", uuidv7.String(), userID, content).
		ExecContext(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to create posts record")
	}

	return uuidv7.String(), nil
}

func (p postsTable) LoadByUUIDTx(ctx context.Context, tx dbr.SessionRunner, uuid string) (*domain.Post, error) {
	m := &domain.Post{}
	err := tx.Select("BIN_TO_UUID(uuid) AS uuid, user_id, content, created_at, updated_at").
		From("posts").
		Where("uuid = UUID_TO_BIN(?) ", uuid).
		LoadOneContext(ctx, m)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return m, nil
		}
		return nil, errors.Wrap(err, "failed to load post records by uuid")
	}

	return m, nil
}
