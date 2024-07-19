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

func Test_userUsecase_SearchUser(t *testing.T) {
	userID := int64(1)
	ctx := appcontext.WithUserID(context.Background(), userID)

	type fields struct {
		userRepo   func(ctrl *gomock.Controller) rdb.Users
		followRepo func(ctrl *gomock.Controller) rdb.Follows
	}
	type args struct {
		ctx context.Context
		req *request.SearchUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.SearchUserResponse
		wantErr bool
	}{
		{
			name: "ユーザー検索結果を取得しレスポンスを返す",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().FindByUsername(ctx, []int64{userID}, "test_username").Return(
						&[]domain.User{
							{ID: 2, Username: "test_username_2", AccountID: "test_account_id_2"},
							{ID: 3, Username: "test_username_3", AccountID: "test_account_id_3"},
							{ID: 4, Username: "test_username_4", AccountID: "test_account_id_4"},
						},
						nil,
					)
					return m
				},
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDs(ctx, userID, []int64{2, 3, 4}).Return(
						&[]domain.Follow{
							{UserID: 1, FollowUserID: 2},
							{UserID: 1, FollowUserID: 3},
						},
						nil,
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SearchUserRequest{Username: "test_username"},
			},
			want: &response.SearchUserResponse{
				Count: 3,
				Users: []response.UserResponse{
					{Username: "test_username_2", AccountID: "test_account_id_2", IsFollowed: true},
					{Username: "test_username_3", AccountID: "test_account_id_3", IsFollowed: true},
					{Username: "test_username_4", AccountID: "test_account_id_4", IsFollowed: false},
				},
			},
		},
		{
			name: "ユーザーのフォロー情報取得時にエラー",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().FindByUsername(ctx, []int64{userID}, "test_username").Return(
						&[]domain.User{
							{ID: 2, Username: "test_username_2", AccountID: "test_account_id_2"},
							{ID: 3, Username: "test_username_3", AccountID: "test_account_id_3"},
							{ID: 4, Username: "test_username_4", AccountID: "test_account_id_4"},
						},
						nil,
					)
					return m
				},
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					m.EXPECT().LoadByUserIDAndFollowUserIDs(ctx, userID, []int64{2, 3, 4}).Return(
						nil,
						fmt.Errorf("error"),
					)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SearchUserRequest{Username: "test_username"},
			},
			wantErr: true,
		},
		{
			name: "検索時にエラー発生",
			fields: fields{
				userRepo: func(ctrl *gomock.Controller) rdb.Users {
					m := mock_rdb.NewMockUsers(ctrl)
					m.EXPECT().FindByUsername(ctx, []int64{userID}, "test_username").Return(nil, fmt.Errorf("error"))
					return m
				},
				followRepo: func(ctrl *gomock.Controller) rdb.Follows {
					m := mock_rdb.NewMockFollows(ctrl)
					return m
				},
			},
			args: args{
				ctx: ctx,
				req: &request.SearchUserRequest{Username: "test_username"},
			},
			wantErr: true,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUsecase{
				userRepo:   tt.fields.userRepo(ctrl),
				followRepo: tt.fields.followRepo(ctrl),
			}
			got, err := u.SearchUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.SearchUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
