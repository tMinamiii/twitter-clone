package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/response"
	"tMinamiii/Tweet/testutil"
	"tMinamiii/Tweet/usecase"
	mock_usecase "tMinamiii/Tweet/usecase/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func Test_followHandler_FollowUser(t *testing.T) {
	type fields struct {
		followUsecase func(ctrl *gomock.Controller) usecase.Follow
	}
	type args struct {
		req request.FollowUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "ユーザーフォローに成功し、結果のJsonを受け取る",
			fields: fields{
				followUsecase: func(ctrl *gomock.Controller) usecase.Follow {
					m := mock_usecase.NewMockFollow(ctrl)
					req := &request.FollowUserRequest{AccountID: "test_account_id"}
					resp := response.NewFollowResponse("test_username", "test_account_id")
					m.EXPECT().FollowUser(gomock.Any(), req).Return(resp, nil)
					return m
				},
			},
			args: args{
				req: request.FollowUserRequest{AccountID: "test_account_id"},
			},
			want: `{
  "username": "test_username",
  "accountId": "test_account_id"
}
`,
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &followHandler{
				followUsecase: tt.fields.followUsecase(ctrl),
			}

			c, rec := testutil.CreateContext(http.MethodPost, "http://localhost:1323/v1/follow", tt.args.req)
			if err := u.FollowUser(c); (err != nil) != tt.wantErr {
				t.Errorf("followHandler.FollowUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var out bytes.Buffer
			json.Indent(&out, rec.Body.Bytes(), "", "  ")
			got := out.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("got diff: (-got +want)\n%s", diff)
			}

		})
	}
}

func Test_followHandler_UnFollowUser(t *testing.T) {
	type fields struct {
		followUsecase func(ctrl *gomock.Controller) usecase.Follow
	}
	type args struct {
		req url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "フォロー解除に成功し、空の結果を受け取る",
			fields: fields{
				followUsecase: func(ctrl *gomock.Controller) usecase.Follow {
					m := mock_usecase.NewMockFollow(ctrl)
					req := &request.UnFollowUserRequest{AccountID: "test_account_id"}
					m.EXPECT().UnFollowUser(gomock.Any(), req).Return(nil)
					return m
				},
			},
			args: args{
				req: url.Values{"accountId": []string{"test_account_id"}},
			},
			want: `{}
`,
		},
	}
	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &followHandler{
				followUsecase: tt.fields.followUsecase(ctrl),
			}
			url := fmt.Sprintf("http://localhost:1323/v1/follow?%s", tt.args.req.Encode())
			c, rec := testutil.CreateContext(http.MethodDelete, url, nil)
			if err := u.UnFollowUser(c); (err != nil) != tt.wantErr {
				t.Errorf("followHandler.UnFollowUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := rec.Body.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("got diff: (-got +want)\n%s", diff)
			}
		})
	}
}
