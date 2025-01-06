package response

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func SendSuccess(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendError(c echo.Context, statusCode int, errorMessage string) error {
	return c.JSON(statusCode, ErrorResponse{
		Error: errorMessage,
	})
}
