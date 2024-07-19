package handler

import (
	"net/http"
	"tMinamiii/Tweet/env"
	"tMinamiii/Tweet/session"

	"github.com/labstack/echo/v4"
)

type Auth interface {
	DummySession(c echo.Context) error
}

type authHandler struct {
}

func NewAuthHandler() Auth {
	return &authHandler{}
}

func (a *authHandler) DummySession(c echo.Context) error {
	if _, err := session.CreateSession(c, env.DummySessionAccountID()); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{})
}
