package usecase

import (
	"context"
	"fmt"
	"reflect"
	"tMinamiii/Tweet/appcontext"
	"tMinamiii/Tweet/domain"
	"tMinamiii/Tweet/infra/rdb"
	mock_rdb "tMinamiii/Tweet/infra/rdb/mock"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/response"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestPostUsecase_Post(t *testing.T) {
	uuid := "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f"
	userID := int64(1)
	content := "tweet"
	username := "test_username"
	accountID := "account_id"
	postedAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	ctx := appcontext.WithUserID(context.Background(), userID)

	type fields struct {
		postRepository func(ctrl *gomock.Controller) rdb.Posts
		userRepository func(ctrl *gomock.Controller) rdb.Users
	}
	type args struct {
		ctx context.Context
		req *request.SubmitPostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.SubmitPostResponse
		wantErr bool
	}{
		{
			name: "記事投稿に成功し、投稿内容のレスポンスが返ってくる",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)

					m.EXPECT().CreateTx(ctx, gomock.Any(), userID, content).Return(uuid, nil)
					m.EXPECT().LoadByUUIDTx(ctx, gomock.Any(), uuid).Return(
						&domain.Post{
							UUID:      uuid,
							UserID:    userID,
							Content:   content,
							CreatedAt: postedAt,
						},
						nil,
					)

					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)

					m.EXPECT().LoadByID(ctx, userID).Return(
						&domain.User{
							ID:        userID,
							AccountID: accountID,
							Username:  username,
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SubmitPostRequest{Content: content},
			},
			want: &response.SubmitPostResponse{
				UUID:      uuid,
				Username:  username,
				AccountID: accountID,
				Content:   content,
				PostedAt:  postedAt,
			},
		},
		{
			name: "記事投稿に成功し、投稿内容のレスポンスが返ってくる",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)

					m.EXPECT().CreateTx(ctx, gomock.Any(), userID, content).Return(uuid, nil)
					m.EXPECT().LoadByUUIDTx(ctx, gomock.Any(), uuid).Return(
						&domain.Post{
							UUID:      uuid,
							UserID:    userID,
							Content:   content,
							CreatedAt: postedAt,
						},
						nil,
					)

					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)

					m.EXPECT().LoadByID(ctx, userID).Return(
						&domain.User{
							ID:        userID,
							AccountID: accountID,
							Username:  username,
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SubmitPostRequest{Content: content},
			},
			want: &response.SubmitPostResponse{
				UUID:      uuid,
				Username:  username,
				AccountID: accountID,
				Content:   content,
				PostedAt:  postedAt,
			},
		},
		{
			name: "Context内にuserIDがないためエラー",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)
					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					return m
				},
			},
			args: args{
				ctx: context.Background(),
				req: &request.SubmitPostRequest{Content: content},
			},
			wantErr: true,
		},
		{
			name: "ユーザーレコード取得時にエラー",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)
					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)

					m.EXPECT().LoadByID(ctx, userID).Return(
						nil,
						fmt.Errorf("error"),
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SubmitPostRequest{Content: content},
			},
			wantErr: true,
		},
		{
			name: "ユーザーレコードが見つからずエラー",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)
					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)

					m.EXPECT().LoadByID(ctx, userID).Return(
						&domain.User{},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SubmitPostRequest{Content: content},
			},
			wantErr: true,
		},
		{
			name: "ポストレコード作成時にエラー",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)

					m.EXPECT().CreateTx(ctx, gomock.Any(), userID, content).Return("", fmt.Errorf("error"))
					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)

					m.EXPECT().LoadByID(ctx, userID).Return(
						&domain.User{
							ID:        userID,
							AccountID: accountID,
							Username:  username,
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SubmitPostRequest{Content: content},
			},
			wantErr: true,
		},
		{
			name: "CreateTx後のポスト取得でエラー",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)

					m.EXPECT().CreateTx(ctx, gomock.Any(), userID, content).Return(uuid, nil)
					m.EXPECT().LoadByUUIDTx(ctx, gomock.Any(), uuid).Return(
						nil,
						fmt.Errorf("error"),
					)

					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)

					m.EXPECT().LoadByID(ctx, userID).Return(
						&domain.User{
							ID:        userID,
							AccountID: accountID,
							Username:  username,
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SubmitPostRequest{Content: content},
			},
			wantErr: true,
		},
		{
			name: "CreateTx後のポスト取得結果が空",
			fields: fields{
				postRepository: func(ctrl *gomock.Controller) rdb.Posts {
					m := mock_rdb.NewMockPosts(ctrl)

					m.EXPECT().CreateTx(ctx, gomock.Any(), userID, content).Return(uuid, nil)
					m.EXPECT().LoadByUUIDTx(ctx, gomock.Any(), uuid).Return(
						&domain.Post{},
						nil,
					)

					return m
				},
				userRepository: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)

					m.EXPECT().LoadByID(ctx, userID).Return(
						&domain.User{
							ID:        userID,
							AccountID: accountID,
							Username:  username,
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SubmitPostRequest{Content: content},
			},
			wantErr: true,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &postUsecase{
				postRepository: tt.fields.postRepository(ctrl),
				userRepository: tt.fields.userRepository(ctrl),
			}
			got, err := p.SubmitPost(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostUsecase.Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostUsecase.Post() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postUsecase_Timeline(t *testing.T) {
	userID := int64(1)
	ctx := appcontext.WithUserID(context.Background(), userID)

	type fields struct {
		postUserRepository func(ctrl *gomock.Controller) rdb.PostsUsers
		followRepository   func(ctrl *gomock.Controller) rdb.Follows
	}
	type args struct {
		ctx context.Context
		req *request.TimelineRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.TimelineResponse
		wantErr bool
	}{
		{
			name: "タイムライン4件取得",
			fields: fields{
				postUserRepository: func(ctrl *gomock.Controller) rdb.PostsUsers {
					m := mock_rdb.NewMockPostsUsers(ctrl)
					uuid := "0189f7e1-ae2c-7809-8aeb-b819cf5e9e7f"
					m.EXPECT().LoadByUserIDs(ctx, []int64{2, 3, 1}, 5, &uuid).Return(
						&[]domain.PostUser{
							{UUID: "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f", UserID: 2, AccountID: "test_account_id_2", Username: "test_username_2", Content: "tweet 1", CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
							{UUID: "0189f7eb-ae2c-7809-8aeb-b819cf5e9e7f", UserID: 3, AccountID: "test_account_id_3", Username: "test_username_3", Content: "tweet 2", CreatedAt: time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC)},
							{UUID: "0189f7ec-ae2c-7809-8aeb-b819cf5e9e7f", UserID: 2, AccountID: "test_account_id_2", Username: "test_username_2", Content: "tweet 3", CreatedAt: time.Date(2024, 1, 1, 2, 0, 0, 0, time.UTC)},
							{UUID: "0189f7ed-ae2c-7809-8aeb-b819cf5e9e7f", UserID: 3, AccountID: "test_account_id_3", Username: "test_username_3", Content: "tweet 4", CreatedAt: time.Date(2024, 1, 1, 3, 0, 0, 0, time.UTC)},
							{UUID: "0189f7ee-ae2c-7809-8aeb-b819cf5e9e7f", UserID: 1, AccountID: "test_account_id_1", Username: "test_username_1", Content: "tweet 5", CreatedAt: time.Date(2024, 1, 1, 4, 0, 0, 0, time.UTC)},
						},
						nil,
					)
					return m
				},
				followRepository: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserID(ctx, userID).Return(
						&[]domain.Follow{
							{ID: 1, UserID: userID, FollowUserID: 2},
							{ID: 1, UserID: userID, FollowUserID: 3},
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.TimelineRequest{
					Limit: 5,
					SinceUUID: func() *string {
						uuid := "0189f7e1-ae2c-7809-8aeb-b819cf5e9e7f"
						return &uuid
					}(),
				},
			},
			want: func() *response.TimelineResponse {
				lastUUID := "0189f7ee-ae2c-7809-8aeb-b819cf5e9e7f"
				return &response.TimelineResponse{
					Count:    5,
					LastUUID: &lastUUID,
					Posts: []response.PostResponse{
						{UUID: "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f", Username: "test_username_2", AccountID: "test_account_id_2", Content: "tweet 1", PostedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
						{UUID: "0189f7eb-ae2c-7809-8aeb-b819cf5e9e7f", Username: "test_username_3", AccountID: "test_account_id_3", Content: "tweet 2", PostedAt: time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC)},
						{UUID: "0189f7ec-ae2c-7809-8aeb-b819cf5e9e7f", Username: "test_username_2", AccountID: "test_account_id_2", Content: "tweet 3", PostedAt: time.Date(2024, 1, 1, 2, 0, 0, 0, time.UTC)},
						{UUID: "0189f7ed-ae2c-7809-8aeb-b819cf5e9e7f", Username: "test_username_3", AccountID: "test_account_id_3", Content: "tweet 4", PostedAt: time.Date(2024, 1, 1, 3, 0, 0, 0, time.UTC)},
						{UUID: "0189f7ee-ae2c-7809-8aeb-b819cf5e9e7f", Username: "test_username_1", AccountID: "test_account_id_1", Content: "tweet 5", PostedAt: time.Date(2024, 1, 1, 4, 0, 0, 0, time.UTC)},
					},
				}
			}(),
		},
		{
			name: "タイムライン0件取得",
			fields: fields{
				postUserRepository: func(ctrl *gomock.Controller) rdb.PostsUsers {
					m := mock_rdb.NewMockPostsUsers(ctrl)
					uuid := "0189f7e1-ae2c-7809-8aeb-b819cf5e9e7f"
					m.EXPECT().LoadByUserIDs(ctx, []int64{2, 3, 1}, 4, &uuid).Return(
						&[]domain.PostUser{},
						nil,
					)
					return m
				},
				followRepository: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserID(ctx, userID).Return(
						&[]domain.Follow{
							{ID: 1, UserID: userID, FollowUserID: 2},
							{ID: 1, UserID: userID, FollowUserID: 3},
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.TimelineRequest{
					Limit: 4,
					SinceUUID: func() *string {
						uuid := "0189f7e1-ae2c-7809-8aeb-b819cf5e9e7f"
						return &uuid
					}(),
				},
			},
			want: func() *response.TimelineResponse {
				return &response.TimelineResponse{
					Count:    0,
					LastUUID: nil,
					Posts:    []response.PostResponse{},
				}
			}(),
		},
		{
			name: "フォローしているユーザー取得時にエラー",
			fields: fields{
				postUserRepository: func(ctrl *gomock.Controller) rdb.PostsUsers {
					m := mock_rdb.NewMockPostsUsers(ctrl)
					return m
				},
				followRepository: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserID(ctx, userID).Return(
						nil,
						fmt.Errorf("error"),
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.TimelineRequest{
					Limit: 4,
					SinceUUID: func() *string {
						uuid := "0189f7e1-ae2c-7809-8aeb-b819cf5e9e7f"
						return &uuid
					}(),
				},
			},
			wantErr: true,
		},
		{
			name: "投稿データ取得時にエラー",
			fields: fields{
				postUserRepository: func(ctrl *gomock.Controller) rdb.PostsUsers {
					m := mock_rdb.NewMockPostsUsers(ctrl)
					uuid := "0189f7e1-ae2c-7809-8aeb-b819cf5e9e7f"
					m.EXPECT().LoadByUserIDs(ctx, []int64{2, 3, 1}, 4, &uuid).Return(
						nil,
						fmt.Errorf("error"),
					)
					return m
				},
				followRepository: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserID(ctx, userID).Return(
						&[]domain.Follow{
							{ID: 1, UserID: userID, FollowUserID: 2},
							{ID: 1, UserID: userID, FollowUserID: 3},
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.TimelineRequest{
					Limit: 4,
					SinceUUID: func() *string {
						uuid := "0189f7e1-ae2c-7809-8aeb-b819cf5e9e7f"
						return &uuid
					}(),
				},
			},
			wantErr: true,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &postUsecase{
				postUserRepository: tt.fields.postUserRepository(ctrl),
				followRepository:   tt.fields.followRepository(ctrl),
			}
			got, err := p.Timeline(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("postUsecase.Timeline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postUsecase.Timeline() = %v, want %v", got, tt.want)
			}
		})
	}
}
