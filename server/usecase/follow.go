//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"context"
	"fmt"
	"net/http"
	"tMinamiii/Tweet/appcontext"
	"tMinamiii/Tweet/infra/rdb"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/response"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Follow interface {
	FollowUser(ctx context.Context, req *request.FollowUserRequest) (*response.FollowUserResponse, error)
	UnFollowUser(ctx context.Context, req *request.UnFollowUserRequest) error
}

type followUsecase struct {
	followRepo rdb.Follows
	userRepo   rdb.Users
}

func NewFollowUsecase() Follow {
	return &followUsecase{
		followRepo: rdb.NewFollowsTable(),
		userRepo:   rdb.NewUsers(),
	}
}

func (f *followUsecase) FollowUser(ctx context.Context, req *request.FollowUserRequest) (*response.FollowUserResponse, error) {
	userID, err := appcontext.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	followUser, err := f.userRepo.LoadByAccountID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}

	if followUser.ID == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("user not found"))
	}

	if userID == followUser.ID {
		return nil, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("can't follow myself"))
	}

	sess := rdb.GetTweetSession()
	tx, err := sess.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to follow user because begin transaction")
	}
	defer tx.RollbackUnlessCommitted()

	// フォロー済みかどうか確認。フォロー済みならCreateしない
	alreadyFollowed, err := f.followRepo.LoadByUserIDAndFollowUserIDTx(ctx, tx, userID, followUser.ID)
	if err != nil {
		return nil, err
	}

	if alreadyFollowed.ID != 0 {
		return response.NewFollowResponse(followUser.Username, followUser.AccountID), nil
	}

	_, err = f.followRepo.CreateTx(ctx, tx, userID, followUser.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to follow user because commit transaction")
	}

	return response.NewFollowResponse(followUser.Username, followUser.AccountID), nil
}

func (f *followUsecase) UnFollowUser(ctx context.Context, req *request.UnFollowUserRequest) error {
	userID, err := appcontext.GetUserID(ctx)
	if err != nil {
		return err
	}

	followUser, err := f.userRepo.LoadByAccountID(ctx, req.AccountID)
	if err != nil {
		return err
	}

	if followUser.ID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("user not found"))
	}

	if userID == followUser.ID {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("can't un-follow myself"))
	}

	sess := rdb.GetTweetSession()
	tx, err := sess.Begin()
	if err != nil {
		return errors.Wrapf(err, "failed to un-follow user because begin transaction")
	}
	defer tx.RollbackUnlessCommitted()

	// フォロー解除済みかどうか確認。解除済みならDeleteしない
	alreadyUnFollowed, err := f.followRepo.LoadByUserIDAndFollowUserIDTx(ctx, tx, userID, followUser.ID)
	if err != nil {
		return err
	}

	if alreadyUnFollowed.ID == 0 {
		return nil
	}

	err = f.followRepo.DeleteTx(ctx, tx, userID, followUser.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrapf(err, "failed to follow user because commit transaction")
	}

	return nil
}
