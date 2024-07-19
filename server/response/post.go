package response

import (
	"time"
)

type SubmitPostResponse struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	AccountID string    `json:"accountId"`
	Content   string    `json:"content"`
	PostedAt  time.Time `json:"postedAt"`
}

func NewSubmitPostResponse(uuid, username, accountID, content string, postedAt time.Time) *SubmitPostResponse {
	return &SubmitPostResponse{
		UUID:      uuid,
		Username:  username,
		AccountID: accountID,
		Content:   content,
		PostedAt:  postedAt,
	}
}

type PostResponse struct {
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	AccountID string    `json:"accountId"`
	Content   string    `json:"content"`
	PostedAt  time.Time `json:"postedAt"`
}

func NewPostResponse(uuid, username, accountID, content string, postedAt time.Time) *PostResponse {
	return &PostResponse{
		UUID:      uuid,
		Username:  username,
		AccountID: accountID,
		Content:   content,
		PostedAt:  postedAt,
	}
}

type TimelineResponse struct {
	Count    int            `json:"count"`
	LastUUID *string        `json:"lastUuid"`
	Posts    []PostResponse `json:"posts"`
}

func NewTimelineResponse(count int, lastUUID *string, posts []PostResponse) *TimelineResponse {
	return &TimelineResponse{
		Count:    count,
		LastUUID: lastUUID,
		Posts:    posts,
	}
}
