package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func CreateContext(method, url string, body interface{}) (echo.Context, *httptest.ResponseRecorder) {
	var req *http.Request
	if method == echo.GET || method == echo.DELETE {
		req, _ = http.NewRequest(method, url, nil)
	} else {
		f, _ := json.Marshal(body)
		req, _ = http.NewRequest(method, url, bytes.NewReader(f))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}

	rec := httptest.NewRecorder()
	return echo.New().NewContext(req, rec), rec
}
