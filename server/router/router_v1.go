package router

import (
	"tMinamiii/Tweet/handler"
	"tMinamiii/Tweet/middleware"

	"github.com/labstack/echo/v4"
)

func RouterV1(e *echo.Echo) {
	v1 := e.Group("v1")
	v1.Use(middleware.APIKey)

	auth := handler.NewAuthHandler()
	v1.GET("/dummy-session", auth.DummySession)

	// after auth
	v1.Use(middleware.SessionAuth)
	user := handler.NewUserHandler()
	v1.GET("/users/search", user.SearchUser)

	post := handler.NewPostHandler()
	v1.GET("/posts/timeline", post.Timeline)
	v1.POST("/posts", post.SubmitPost)

	follow := handler.NewFollowHandler()
	v1.POST("/follows", follow.FollowUser)
	v1.DELETE("/follows", follow.UnFollowUser)
}
