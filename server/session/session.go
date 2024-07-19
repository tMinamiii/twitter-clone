package session

import (
	"fmt"
	"tMinamiii/Tweet/env"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	accountIDKey = "accountId"
)

func CreateSession(c echo.Context, accountID string) (*sessions.Session, error) {
	sess, err := session.Get(env.SessionName(), c)
	if err != nil {
		return nil, err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values[accountIDKey] = accountID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return nil, err
	}

	return sess, nil
}

func AccountID(c echo.Context) (string, error) {
	sess, err := getSession(c)
	if err != nil {
		return "", nil
	}

	accountIDValue, ok := sess.Values[accountIDKey]
	if !ok {
		err := fmt.Errorf("session value not found")
		return "", err
	}

	accountID, ok := accountIDValue.(string)
	if !ok {
		err := fmt.Errorf("failed to assert session value")
		return "", err
	}

	return accountID, nil
}

func getSession(c echo.Context) (*sessions.Session, error) {
	sess, err := session.Get(env.SessionName(), c)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
