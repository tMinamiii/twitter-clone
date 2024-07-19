//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE

package usecase

import (
	"context"
	"fmt"
	"tMinamiii/Tweet/appcontext"
	"tMinamiii/Tweet/infra/rdb"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/response"

	"github.com/pkg/errors"
)

type Post interface {
	SubmitPost(ctx context.Context, req *request.SubmitPostRequest) (*response.SubmitPostResponse, error)
	Timeline(ctx context.Context, req *request.TimelineRequest) (*response.TimelineResponse, error)
}

type postUsecase struct {
	postRepository     rdb.Posts
	userRepository     rdb.Users
	postUserRepository rdb.PostsUsers
	followRepository   rdb.Follows
}

func NewPostUsecase() Post {
	return &postUsecase{
		postRepository:     rdb.NewPostsTable(),
		userRepository:     rdb.NewUsers(),
		postUserRepository: rdb.NewPostsUsers(),
		followRepository:   rdb.NewFollowsTable(),
	}
}

func (p *postUsecase) SubmitPost(ctx context.Context, req *request.SubmitPostRequest) (*response.SubmitPostResponse, error) {
	userID, err := appcontext.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := p.userRepository.LoadByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("failed to post due to user record not found. userID = %d", userID)
	}

	sess := rdb.GetTweetSession()
	tx, err := sess.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get transaction")
	}
	defer tx.RollbackUnlessCommitted()

	uuid, err := p.postRepository.CreateTx(ctx, tx, userID, req.Content)
	if err != nil {
		return nil, err
	}

	post, err := p.postRepository.LoadByUUIDTx(ctx, tx, uuid)
	if err != nil {
		return nil, err
	}

	if post.UUID == "" {
		return nil, fmt.Errorf("load by uuid tx result is empty")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to commit post")
	}

	return response.NewSubmitPostResponse(post.UUID, user.Username, user.AccountID, post.Content, post.CreatedAt), nil
}

func (p *postUsecase) Timeline(ctx context.Context, req *request.TimelineRequest) (*response.TimelineResponse, error) {
	userID, err := appcontext.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	follows, err := p.followRepository.LoadByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	timelineUserIDs := make([]int64, 0, len(*follows))
	for _, v := range *follows {
		timelineUserIDs = append(timelineUserIDs, v.FollowUserID)
	}

	// タイムラインには自信のつぶやきも表示されるため、自信のuserIDも追加する
	timelineUserIDs = append(timelineUserIDs, userID)

	postsUsers, err := p.postUserRepository.LoadByUserIDs(ctx, timelineUserIDs, req.Limit, req.SinceUUID)
	if err != nil {
		return nil, err
	}

	count := len(*postsUsers)
	postResps := make([]response.PostResponse, 0, count)
	for _, v := range *postsUsers {
		p := response.NewPostResponse(v.UUID, v.Username, v.AccountID, v.Content, v.CreatedAt)
		postResps = append(postResps, *p)
	}
	var lastUUID *string
	if count >= 1 {
		uuid := (*postsUsers)[count-1].UUID
		lastUUID = &uuid
	}

	return response.NewTimelineResponse(count, lastUUID, postResps), nil
}
