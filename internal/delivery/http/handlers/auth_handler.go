package handlers

import (
	"net/http"

	"github.com/hutamy/golang-clean-architecture/internal/domain"
	"github.com/hutamy/golang-clean-architecture/pkg/response"
	"github.com/hutamy/golang-clean-architecture/pkg/validator"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var user validator.UserLogin
	if err := c.Bind(&user); err != nil {
		return response.SendError(c, http.StatusBadRequest, "Invalid payload")
	}

	if err := validator.ValidateLogin(user); err != nil {
		return response.SendError(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.authService.Login(c.Request().Context(), user.Username, user.Password)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, err.Error())
	}

	return response.SendSuccess(c, http.StatusOK, "Login Success", map[string]interface{}{
		"token": token,
	})
}
