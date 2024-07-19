package handler

import (
	"net/http"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/usecase"

	"github.com/labstack/echo/v4"
)

type User interface {
	SearchUser(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.User
}

func NewUserHandler() User {
	return &userHandler{
		userUsecase: usecase.NewUserUsecase(),
	}
}

func (u *userHandler) SearchUser(c echo.Context) error {
	req, err := request.NewSearchUserRequest(c)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := u.userUsecase.SearchUser(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
