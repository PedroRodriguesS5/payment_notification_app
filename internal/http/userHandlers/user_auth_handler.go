package userHandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/internal/middleware"
	"github.com/pedroRodriguesS5/payment_notification/internal/service/user"
)

func RegisterUserAuthRoutes(e *echo.Echo, uService user.Service) {
	authGroup := e.Group("/auth")

	authGroup.Use(middleware.AuthToken)
	authGroup.GET("/user/:id", GetUser(uService))
}

func GetUser(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		findedUser, err := s.GetUser(c.Request().Context(), userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found"})
		}
		if findedUser == nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, findedUser)
	}
}

// func GetAllUsers(s user.Service) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		allUsers, err := s.GetAllUsers(c.Request().Context())

// 		if err != nil {
// 			return c.String(http.StatusInternalServerError, err.Error())
// 		}

// 		return c.JSON(http.StatusOK, allUsers)
// 	}
// }
