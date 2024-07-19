package response

type UserResponse struct {
	Username   string `json:"username"`
	AccountID  string `json:"accountId"`
	IsFollowed bool   `json:"isFollowed"`
}

func NewUserResponse(username, accountID string, isFollowed bool) *UserResponse {
	return &UserResponse{
		Username:   username,
		AccountID:  accountID,
		IsFollowed: isFollowed,
	}
}

type SearchUserResponse struct {
	Count int            `json:"count"`
	Users []UserResponse `json:"users"`
}

func NewSearchUserResponse(count int, users []UserResponse) *SearchUserResponse {
	return &SearchUserResponse{
		Count: count,
		Users: users,
	}
}
