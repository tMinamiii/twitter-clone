package request

import (
	"fmt"
	"net/http"
	"reflect"
	"tMinamiii/Tweet/testutil"
	"testing"
)

func TestNewSearchUserRequest(t *testing.T) {
	tests := []struct {
		name    string
		params  string
		want    *SearchUserRequest
		wantErr bool
	}{
		{
			name:   "リクエストのBindに成功",
			params: "?username=test_username",
			want:   &SearchUserRequest{Username: "test_username"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:1323/v1/users/search%s", tt.params)
			c, _ := testutil.CreateContext(http.MethodGet, url, nil)
			got, err := NewSearchUserRequest(c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSearchUserRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchUserRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
