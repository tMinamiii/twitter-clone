package response

import (
	"reflect"
	"testing"
	"time"
)

func TestNewPostResponse(t *testing.T) {
	uuid := "0000-0000-0000-0000"
	username := "test_username"
	accountID := "test_account_id"
	content := "tweet"
	postedAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	want := &SubmitPostResponse{
		UUID:      uuid,
		Username:  username,
		AccountID: accountID,
		Content:   content,
		PostedAt:  postedAt,
	}

	if got := NewSubmitPostResponse(uuid, username, accountID, content, postedAt); !reflect.DeepEqual(got, want) {
		t.Errorf("NewPostResponse() = %v, want %v", got, want)
	}
}

func TestNewTimelineResponse(t *testing.T) {
	posts := []PostResponse{
		{
			UUID:      "0000-0000-0000-0001",
			Username:  "test_username_1",
			AccountID: "test_account_id_1",
			Content:   "tweet 1",
			PostedAt:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			UUID:      "0000-0000-0000-0002",
			Username:  "test_username_2",
			AccountID: "test_account_id_2",
			Content:   "tweet 2",
			PostedAt:  time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}
	count := 2
	lastUUID := "0000-0000-0000-0002"
	want := &TimelineResponse{
		Count:    count,
		LastUUID: &lastUUID,
		Posts:    posts,
	}
	if got := NewTimelineResponse(count, &lastUUID, posts); !reflect.DeepEqual(got, want) {
		t.Errorf("NewTimelineResponse() = %v, want %v", got, want)
	}
}
