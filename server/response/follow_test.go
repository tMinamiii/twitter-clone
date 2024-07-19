package response

import (
	"reflect"
	"testing"
)

func TestNewFollowResponse(t *testing.T) {
	username := "test_username"
	accountID := "test_account_id"
	want := &FollowUserResponse{
		Username:  username,
		AccountID: accountID,
	}
	if got := NewFollowResponse(username, accountID); !reflect.DeepEqual(got, want) {
		t.Errorf("NewFollowResponse() = %v, want %v", got, want)
	}
}
