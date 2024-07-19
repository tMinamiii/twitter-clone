package response

type FollowUserResponse struct {
	Username  string `json:"username"`
	AccountID string `json:"accountId"`
}

func NewFollowResponse(username, accountID string) *FollowUserResponse {
	return &FollowUserResponse{
		Username:  username,
		AccountID: accountID,
	}
}
