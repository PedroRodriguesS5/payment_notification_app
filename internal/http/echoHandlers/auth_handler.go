package echoHandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/internal/user"
)

func RegisterAuthRoutes(e *echo.Echo, uService user.Service) {
	authGroup := e.Group("/auth")
	authGroup.GET("/user/:id", GetUser(uService))
}

func GetUser(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		findedUser, err := s.GetUser(c.Request().Context(), userId)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if findedUser == nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, findedUser)
	}
}
