package handler

import (
	"net/http"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/usecase"

	"github.com/labstack/echo/v4"
)

type Follow interface {
	FollowUser(c echo.Context) error
	UnFollowUser(c echo.Context) error
}

type followHandler struct {
	followUsecase usecase.Follow
}

func NewFollowHandler() Follow {
	return &followHandler{
		followUsecase: usecase.NewFollowUsecase(),
	}
}

func (u *followHandler) FollowUser(c echo.Context) error {
	req, err := request.NewFollowUserRequest(c)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := u.followUsecase.FollowUser(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, resp)
}

func (u *followHandler) UnFollowUser(c echo.Context) error {
	req, err := request.NewUnFollowUserRequest(c)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	err = u.followUsecase.UnFollowUser(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{})
}
