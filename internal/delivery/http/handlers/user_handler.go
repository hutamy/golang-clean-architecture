package handlers

import (
	"net/http"
	"strconv"

	"github.com/hutamy/golang-clean-architecture/internal/domain"
	"github.com/hutamy/golang-clean-architecture/internal/domain/entity"
	"github.com/hutamy/golang-clean-architecture/pkg/response"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return response.SendError(c, http.StatusBadRequest, "Invalid payload")
	}

	if err := h.userService.Register(c.Request().Context(), user); err != nil {
		return response.SendError(c, http.StatusInternalServerError, err.Error())
	}

	return response.SendSuccess(c, http.StatusCreated, "Register Success", nil)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SendError(c, http.StatusBadRequest, "Invalid payload")
	}

	user, err := h.userService.GetUser(c.Request().Context(), id)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, err.Error())
	}

	return response.SendSuccess(c, http.StatusOK, "Success", user)
}
