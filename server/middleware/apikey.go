package middleware

import (
	"log"
	"net/http"
	"tMinamiii/Tweet/env"

	"github.com/labstack/echo/v4"
)

// APIKey ヘッダーのApiKeyを読みこんで一致しているか検証する
func APIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h := c.Request().Header
		apiKey := h.Get("X-API-Key")
		if apiKey == "" {
			log.Println("there is no api key in header")
			return echo.NewHTTPError(http.StatusUnauthorized, "no api key")
		}

		if apiKey != env.APIKey() {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid api key")
		}

		return next(c)
	}
}
