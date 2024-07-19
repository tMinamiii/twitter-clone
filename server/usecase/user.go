//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"context"
	"strings"
	"tMinamiii/Tweet/appcontext"
	"tMinamiii/Tweet/infra/rdb"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/response"
)

type User interface {
	SearchUser(ctx context.Context, req *request.SearchUserRequest) (*response.SearchUserResponse, error)
}

type userUsecase struct {
	userRepo   rdb.Users
	followRepo rdb.Follows
}

func NewUserUsecase() User {
	return &userUsecase{
		userRepo:   rdb.NewUsers(),
		followRepo: rdb.NewFollowsTable(),
	}
}

func (u *userUsecase) SearchUser(ctx context.Context, req *request.SearchUserRequest) (*response.SearchUserResponse, error) {
	userID, err := appcontext.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	trimUsername := strings.TrimSpace(req.Username)
	users, err := u.userRepo.FindByUsername(ctx, []int64{userID}, trimUsername)
	if err != nil {
		return nil, err
	}

	resultUserIDs := make([]int64, 0, len(*users))
	for _, v := range *users {
		resultUserIDs = append(resultUserIDs, v.ID)
	}

	follows, err := u.followRepo.LoadByUserIDAndFollowUserIDs(ctx, userID, resultUserIDs)
	if err != nil {
		return nil, err
	}

	followUserIDSet := make(map[int64]struct{}, len(*follows))
	for _, v := range *follows {
		followUserIDSet[v.FollowUserID] = struct{}{}
	}

	count := len(*users)
	userResps := make([]response.UserResponse, 0, count)
	for _, v := range *users {
		_, isFollowed := followUserIDSet[v.ID]
		r := response.NewUserResponse(v.Username, v.AccountID, isFollowed)
		userResps = append(userResps, *r)
	}

	return response.NewSearchUserResponse(count, userResps), nil
}
