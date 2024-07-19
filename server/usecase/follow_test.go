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

	"github.com/golang/mock/gomock"
)

func Test_followUsecase_FollowUser(t *testing.T) {
	userID := int64(1)
	ctx := appcontext.WithUserID(context.Background(), userID)

	type fields struct {
		followRepo func(ctrl *gomock.Controller) rdb.Follows
		userRepo   func(ctrl *gomock.Controller) rdb.Users
	}
	type args struct {
		ctx context.Context
		req *request.FollowUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.FollowUserResponse
		wantErr bool
	}{
		{
			name: "ユーザーをフォロー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, int64(2)).
						Return(&domain.Follow{}, nil)
					m.EXPECT().CreateTx(ctx, gomock.Any(), userID, int64(2)).Return(int64(1), nil)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			want: &response.FollowUserResponse{
				Username:  "test_username",
				AccountID: "test_account_id",
			},
		},
		{
			name: "すでにフォロー済みだった場合",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, int64(2)).
						Return(&domain.Follow{ID: 1, UserID: userID, FollowUserID: int64(2)}, nil)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			want: &response.FollowUserResponse{
				Username:  "test_username",
				AccountID: "test_account_id",
			},
		},
		{
			name: "自分自身をフォローしようとした場合",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: userID, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "AccountIDでユーザーレコード取得時にエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						nil,
						fmt.Errorf("error"),
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "フォローしようとしたユーザーが不在だった",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(&domain.User{}, nil)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "フォローレコードの存在確時にエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, int64(2)).
						Return(nil, fmt.Errorf("error"))
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "フォローレコード作成時にエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, int64(2)).
						Return(&domain.Follow{}, nil)
					m.EXPECT().CreateTx(ctx, gomock.Any(), userID, int64(2)).Return(int64(0), fmt.Errorf("error"))
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "Context内にuserIDがないためエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					return m
				},
			},
			args: args{
				ctx: context.Background(),
				req: &request.FollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &followUsecase{
				followRepo: tt.fields.followRepo(ctrl),
				userRepo:   tt.fields.userRepo(ctrl),
			}
			got, err := f.FollowUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("followUsecase.FollowUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("followUsecase.FollowUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_followUsecase_UnFollowUser(t *testing.T) {
	userID := int64(1)
	ctx := appcontext.WithUserID(context.Background(), userID)

	type fields struct {
		followRepo func(ctrl *gomock.Controller) rdb.Follows
		userRepo   func(ctrl *gomock.Controller) rdb.Users
	}
	type args struct {
		ctx context.Context
		req *request.UnFollowUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ユーザーをフォロー解除",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					followUserID := int64(2)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, followUserID).
						Return(&domain.Follow{ID: 1, UserID: userID, FollowUserID: followUserID}, nil)
					m.EXPECT().DeleteTx(ctx, gomock.Any(), userID, followUserID).Return(nil)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.UnFollowUserRequest{AccountID: "test_account_id"},
			},
		},
		{
			name: "すでにフォロー解除済みだった場合",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, int64(2)).
						Return(&domain.Follow{}, nil)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.UnFollowUserRequest{AccountID: "test_account_id"},
			},
		},
		{
			name: "フォローレコード削除にエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					followUserID := int64(2)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, followUserID).
						Return(&domain.Follow{ID: 1, UserID: userID, FollowUserID: followUserID}, nil)
					m.EXPECT().DeleteTx(ctx, gomock.Any(), userID, followUserID).Return(fmt.Errorf("error"))
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.UnFollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "フォロー解除済みの確認時にエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDTx(ctx, gomock.Any(), userID, int64(2)).
						Return(nil, fmt.Errorf("errok"))
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 2, Username: "test_username", AccountID: "test_account_id"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.UnFollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "自分をフォロー解除使用としたらエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(
						&domain.User{ID: 1, Username: "myself", AccountID: "myself"},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.UnFollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "解除するユーザーが存在しなければエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(&domain.User{}, nil)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.UnFollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
		{
			name: "解除するユーザーレコードを取得使用としてエラー",
			fields: fields{
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().LoadByAccountID(ctx, "test_account_id").Return(nil, fmt.Errorf("error"))
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.UnFollowUserRequest{AccountID: "test_account_id"},
			},
			wantErr: true,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &followUsecase{
				followRepo: tt.fields.followRepo(ctrl),
				userRepo:   tt.fields.userRepo(ctrl),
			}
			err := f.UnFollowUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("followUsecase.UnFollowUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
