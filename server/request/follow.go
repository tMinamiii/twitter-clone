package request

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type FollowUserRequest struct {
	AccountID string `json:"accountId"` // フォローしたいユーザーのAccountID
}

func NewFollowUserRequest(c echo.Context) (*FollowUserRequest, error) {
	f := &FollowUserRequest{}
	if err := c.Bind(f); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return f, nil
}

type UnFollowUserRequest struct {
	AccountID string `query:"accountId"` // フォローしたいユーザーのAccountID
}

func NewUnFollowUserRequest(c echo.Context) (*UnFollowUserRequest, error) {
	f := &UnFollowUserRequest{}
	if err := c.Bind(f); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return f, nil
}
