package http

import (
	"net/http"

	"github.com/hutamy/golang-clean-architecture/internal/delivery/http/handlers"
	"github.com/hutamy/golang-clean-architecture/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(
	e *echo.Echo,
	authMiddleware *middleware.AuthMiddleware,
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
) {
	e.POST("/users", userHandler.RegisterUser)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/login", authHandler.Login)
	e.GET("/protected", func(c echo.Context) error {
		userID := c.Get("user_id").(string)
		return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to the protected route", "user_id": userID})
	}, authMiddleware.ValidateJWT)
}
