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
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func Test_postHandler_Timeline(t *testing.T) {
	type fields struct {
		postUsecase func(ctrl *gomock.Controller) usecase.Post
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "タイムライン取得に成功し、結果のJsonを受け取る",
			fields: fields{
				postUsecase: func(ctrl *gomock.Controller) usecase.Post {
					m := mock_usecase.NewMockPost(ctrl)
					uuid := "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f"
					req := &request.TimelineRequest{Limit: 2, SinceUUID: &uuid}
					lastUUID := "0189f7ec-ae2c-7809-8aeb-b819cf5e9e7f"
					resp := response.NewTimelineResponse(2, &lastUUID, []response.PostResponse{
						{UUID: "0189f7eb-ae2c-7809-8aeb-b819cf5e9e7f", Username: "test_username_1", AccountID: "test_account_id_1", Content: "tweet 1", PostedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
						{UUID: "0189f7ec-ae2c-7809-8aeb-b819cf5e9e7f", Username: "test_username_2", AccountID: "test_account_id_2", Content: "tweet 2", PostedAt: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)},
					})
					m.EXPECT().Timeline(gomock.Any(), req).Return(resp, nil)
					return m
				},
			},
			args: args{
				query: func() string {
					q := make(url.Values, 1)
					q.Add("limit", "2")
					q.Add("sinceUuid", "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f")
					return q.Encode()
				}(),
			},
			want: `{
  "count": 2,
  "lastUuid": "0189f7ec-ae2c-7809-8aeb-b819cf5e9e7f",
  "posts": [
    {
      "uuid": "0189f7eb-ae2c-7809-8aeb-b819cf5e9e7f",
      "username": "test_username_1",
      "accountId": "test_account_id_1",
      "content": "tweet 1",
      "postedAt": "2024-01-01T00:00:00Z"
    },
    {
      "uuid": "0189f7ec-ae2c-7809-8aeb-b819cf5e9e7f",
      "username": "test_username_2",
      "accountId": "test_account_id_2",
      "content": "tweet 2",
      "postedAt": "2024-01-02T00:00:00Z"
    }
  ]
}
`,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &postHandler{
				postUsecase: tt.fields.postUsecase(ctrl),
			}
			url := fmt.Sprintf("http://localhost:1323/v1/posts/timeline?%s", tt.args.query)
			c, rec := testutil.CreateContext(http.MethodGet, url, nil)
			if err := u.Timeline(c); (err != nil) != tt.wantErr {
				t.Errorf("postHandler.Timeline() error = %v, wantErr %v", err, tt.wantErr)
			}

			var out bytes.Buffer
			json.Indent(&out, rec.Body.Bytes(), "", "  ")
			got := out.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("got diff: (-got, +want)\n%s", diff)
			}
		})
	}
}

func Test_postHandler_SubmitPost(t *testing.T) {
	type fields struct {
		postUsecase func(ctrl *gomock.Controller) usecase.Post
	}
	type args struct {
		req request.SubmitPostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			name: "投稿に成功し結果のJsonを受け取る",
			fields: fields{
				postUsecase: func(ctrl *gomock.Controller) usecase.Post {
					m := mock_usecase.NewMockPost(ctrl)
					req := &request.SubmitPostRequest{Content: "tweet 1"}
					uuid := "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f"
					resp := response.NewSubmitPostResponse(uuid, "test_username", "test_account_id", "tweet 1", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
					m.EXPECT().SubmitPost(gomock.Any(), req).Return(resp, nil)
					return m
				},
			},
			args: args{
				req: request.SubmitPostRequest{Content: "tweet 1"},
			},
			want: `{
  "uuid": "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f",
  "username": "test_username",
  "accountId": "test_account_id",
  "content": "tweet 1",
  "postedAt": "2024-01-01T00:00:00Z"
}
`,
		},
	}

	ctrl := gomock.NewController(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &postHandler{
				postUsecase: tt.fields.postUsecase(ctrl),
			}

			c, rec := testutil.CreateContext(http.MethodPost, "http://localhost:1323/v1/posts", tt.args.req)
			if err := u.SubmitPost(c); (err != nil) != tt.wantErr {
				t.Errorf("postHandler.SubmitPost() error = %v, wantErr %v", err, tt.wantErr)
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
