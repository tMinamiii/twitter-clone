package request

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SearchUserRequest struct {
	Username string `query:"username"`
}

func NewSearchUserRequest(c echo.Context) (*SearchUserRequest, error) {
	f := &SearchUserRequest{}
	if err := c.Bind(f); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return f, nil
}
