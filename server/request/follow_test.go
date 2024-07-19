package request

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"tMinamiii/Tweet/testutil"
	"testing"
)

func TestNewFollowUserRequest(t *testing.T) {
	tests := []struct {
		name    string
		params  map[string]any
		want    *FollowUserRequest
		wantErr bool
	}{
		{
			name: "リクエストのBindに成功",
			params: map[string]any{
				"accountId": "test_follow_account_id",
			},
			want: &FollowUserRequest{AccountID: "test_follow_account_id"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := testutil.CreateContext(http.MethodPost, "http://localhost:1323/v1/follows", tt.params)
			got, err := NewFollowUserRequest(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFollowUserRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFollowUserRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnFollowUserRequest(t *testing.T) {
	tests := []struct {
		name    string
		params  url.Values
		want    *UnFollowUserRequest
		wantErr bool
	}{
		{
			name: "リクエストのBindに成功",
			params: url.Values{
				"accountId": []string{"test_unfollow_account_id"},
			},
			want: &UnFollowUserRequest{AccountID: "test_unfollow_account_id"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:1323/v1/follows?%s", tt.params.Encode())
			c, _ := testutil.CreateContext(http.MethodDelete, url, nil)
			got, err := NewUnFollowUserRequest(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUnFollowUserRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnFollowUserRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
