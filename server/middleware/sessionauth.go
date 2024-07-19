package middleware

import (
	"tMinamiii/Tweet/appcontext"
	"tMinamiii/Tweet/infra/rdb"
	"tMinamiii/Tweet/session"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID, err := session.AccountID(c)
		if err != nil {
			return errors.Wrap(err, "failed to get account id from session value")
		}

		ctx := c.Request().Context()
		user, err := rdb.NewUsers().LoadByAccountID(ctx, accountID)
		if err != nil {
			return errors.Wrap(err, "failed to auth because user not found")
		}

		newCtx := appcontext.WithUserID(ctx, user.ID)
		req := c.Request().WithContext(newCtx)
		c.SetRequest(req)

		return next(c)
	}
}
