//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package rdb

import (
	"context"
	"fmt"
	"log"
	"tMinamiii/Tweet/domain"

	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type Users interface {
	LoadByID(ctx context.Context, id int64) (*domain.User, error)
	LoadByAccountID(ctx context.Context, accountID string) (*domain.User, error)
	FindByUsername(ctx context.Context, exceptUserIDs []int64, username string) (*[]domain.User, error)
}

type usersTable struct {
}

func NewUsers() Users {
	return &usersTable{}
}

func (u *usersTable) LoadByID(ctx context.Context, id int64) (*domain.User, error) {
	m := &domain.User{}

	err := GetTweetSession().
		Select("*").
		From("users").
		Where("id = ?", id).
		LoadOneContext(ctx, m)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return m, nil
		}
		err = errors.Wrapf(err, "failed to load user record by user id")
		log.Println(err)
		return nil, err
	}

	return m, nil
}

func (u *usersTable) LoadByAccountID(ctx context.Context, accountID string) (*domain.User, error) {
	m := &domain.User{}

	err := GetTweetSession().
		Select("*").
		From("users").
		Where("account_id = ?", accountID).
		LoadOneContext(ctx, m)

	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return m, nil
		}
		err = errors.Wrapf(err, "failed to load user record by account id")
		log.Println(err)
		return nil, err
	}

	return m, nil
}

func (u *usersTable) FindByUsername(ctx context.Context, exceptUserIDs []int64, username string) (*[]domain.User, error) {
	m := &[]domain.User{}

	stmt := GetTweetSession().
		Select("*").
		From("users").
		Where("id NOT IN ?", exceptUserIDs)

	if username != "" {
		wildcardUsername := fmt.Sprintf("%%%s%%", username)
		stmt = stmt.Where("username LIKE ?", wildcardUsername)
	}

	_, err := stmt.LoadContext(ctx, m)
	if err != nil {
		err = errors.Wrapf(err, "failed to find user records by username")
		log.Println(ctx, err)
		return nil, err
	}

	return m, nil
}
