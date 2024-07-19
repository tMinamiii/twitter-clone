package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"tMinamiii/Tweet/env"
	"tMinamiii/Tweet/project"
	"tMinamiii/Tweet/router"
	"time"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if env.Env() == "local" {
		godotenv.Load(project.Root() + "/.env.local")
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.PATCH, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAcceptEncoding,
			"X-API-Key",
		},
	}))

	// 複数台構成にする場合は、セッションの保存場所をRedisやDBに移す
	// https://github.com/gorilla/sessions?tab=readme-ov-file#store-implementations
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	router.RouterV1(e)
	// echo graceful shutdown
	// https://echo.labstack.com/docs/cookbook/graceful-shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
