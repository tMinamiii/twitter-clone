package response

import (
	"reflect"
	"testing"
)

func TestNewUserResponse(t *testing.T) {
	username := "test_username"
	accountID := "test_account_id"
	isFollowed := true
	want := &UserResponse{
		Username:   username,
		AccountID:  accountID,
		IsFollowed: true,
	}
	if got := NewUserResponse(username, accountID, isFollowed); !reflect.DeepEqual(got, want) {
		t.Errorf("NewUserResponse() = %v, want %v", got, want)
	}
}

func TestNewSearchUserResponse(t *testing.T) {
	count := 2
	users := []UserResponse{
		{Username: "test_username_1", AccountID: "test_account_id_1", IsFollowed: true},
		{Username: "test_username_2", AccountID: "test_account_id_2", IsFollowed: false},
	}
	want := &SearchUserResponse{Count: count, Users: users}
	if got := NewSearchUserResponse(count, users); !reflect.DeepEqual(got, want) {
		t.Errorf("NewSearchUserResponse() = %v, want %v", got, want)
	}
}
