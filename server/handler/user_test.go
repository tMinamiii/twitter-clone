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

func Test_userHandler_SearchUser(t *testing.T) {
	type fields struct {
		userUsecase func(ctrl *gomock.Controller) usecase.User
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ユーザー検索を行い、検索結果のJsonを受け取る",
			fields: fields{
				userUsecase: func(ctrl *gomock.Controller) usecase.User {
					m := mock_usecase.NewMockUser(ctrl)
					m.EXPECT().
						SearchUser(gomock.Any(), &request.SearchUserRequest{Username: "test_username"}).
						Return(
							response.NewSearchUserResponse(2, []response.UserResponse{
								{Username: "test_username_1", AccountID: "test_account_id_1", IsFollowed: true},
								{Username: "test_username_2", AccountID: "test_account_id_2", IsFollowed: false},
							}),
							nil,
						)
					return m
				},
			},
			args: args{
				query: func() string {
					q := make(url.Values, 1)
					q.Set("username", "test_username")
					return q.Encode()
				}(),
			},
			want: `{
  "count": 2,
  "users": [
    {
      "username": "test_username_1",
      "accountId": "test_account_id_1",
      "isFollowed": true
    },
    {
      "username": "test_username_2",
      "accountId": "test_account_id_2",
      "isFollowed": false
    }
  ]
}
`,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userHandler{
				userUsecase: tt.fields.userUsecase(ctrl),
			}
			url := fmt.Sprintf("http://localhost:1323/v1/users/search?%s", tt.args.query)
			c, rec := testutil.CreateContext(http.MethodGet, url, nil)
			if err := u.SearchUser(c); (err != nil) != tt.wantErr {
				t.Errorf("userHandler.SearchUser() error = %v, wantErr %v", err, tt.wantErr)
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
