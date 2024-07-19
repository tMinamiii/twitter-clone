package handler

import (
	"net/http"
	"tMinamiii/Tweet/request"
	"tMinamiii/Tweet/usecase"

	"github.com/labstack/echo/v4"
)

type Post interface {
	Timeline(c echo.Context) error
	SubmitPost(c echo.Context) error
}

type postHandler struct {
	postUsecase usecase.Post
}

func NewPostHandler() Post {
	return &postHandler{
		postUsecase: usecase.NewPostUsecase(),
	}
}

func (u *postHandler) Timeline(c echo.Context) error {
	req, err := request.NewTimelineRequest(c)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := u.postUsecase.Timeline(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (u *postHandler) SubmitPost(c echo.Context) error {
	req, err := request.NewSubmitPostRequest(c)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := u.postUsecase.SubmitPost(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, resp)
}
