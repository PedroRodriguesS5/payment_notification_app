package echo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/internal/user"
)

func Handlres(uServcie user.Service) *echo.Echo {
	e := echo.New()
	// usersHandlers
	e.GET("/user/:id", GetUser(uServcie))
	e.GET("/user/all", GetAllUsers(uServcie))
	e.POST("/user/create", CreateUser(uServcie))
	return e
}

func CreateUser(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(user.UserRegisterDTO)
		create, err := s.CreateUser(c.Request().Context(), *u)

		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, create)
	}
}

func GetUser(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		fmt.Println(userId)
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

func GetAllUsers(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		allUsers, err := s.GetAllUsers(c.Request().Context())

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, allUsers)
	}
}
