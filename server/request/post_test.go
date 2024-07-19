package request

import (
	"fmt"
	"net/http"
	"reflect"
	"tMinamiii/Tweet/testutil"
	"testing"
)

func TestNewSubmitPostRequest(t *testing.T) {
	tests := []struct {
		name    string
		params  map[string]any
		want    *SubmitPostRequest
		wantErr bool
	}{
		{
			name: "リクエストのBindに成功",
			params: map[string]any{
				"content": "tweet",
			},
			want: &SubmitPostRequest{Content: "tweet"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := testutil.CreateContext(http.MethodPost, "http://localhost:1323/v1/posts", tt.params)
			got, err := NewSubmitPostRequest(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSubmitPostRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubmitPostRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTimelineRequest(t *testing.T) {
	uuid := "0189f7ea-ae2c-7809-8aeb-b819cf5e9e7f"

	tests := []struct {
		name    string
		params  string
		want    *TimelineRequest
		wantErr bool
	}{
		{
			name:   "リクエストのBindに成功",
			params: fmt.Sprintf("?limit=50&sinceUuid=%s", uuid),
			want:   &TimelineRequest{Limit: 50, SinceUUID: &uuid},
		},
		{
			name:   "sinceUUIDなしのリクエストのBindに成功",
			params: "?limit=50",
			want:   &TimelineRequest{Limit: 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:1323/v1/posts/timeline%s", tt.params)
			c, _ := testutil.CreateContext(http.MethodGet, url, tt.params)
			got, err := NewTimelineRequest(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimelineRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimelineRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
