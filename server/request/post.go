package request

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SubmitPostRequest struct {
	Content string `json:"content"`
}

func NewSubmitPostRequest(c echo.Context) (*SubmitPostRequest, error) {
	f := &SubmitPostRequest{}
	if err := c.Bind(f); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return f, nil
}

type TimelineRequest struct {
	Limit     int     `query:"limit"`
	SinceUUID *string `query:"sinceUuid"`
}

func NewTimelineRequest(c echo.Context) (*TimelineRequest, error) {
	f := &TimelineRequest{}
	if err := c.Bind(f); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return f, nil
}
