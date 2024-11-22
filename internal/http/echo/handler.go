package echo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/internal/user"
	"github.com/pedroRodriguesS5/payment_notification/pkg/utils"
	"github.com/pedroRodriguesS5/payment_notification/pkg/utils/tools"
)

func Handlres(uServcie user.Service) *echo.Echo {
	e := echo.New()
	// usersHandlers
	e.GET("/user/:id", GetUser(uServcie))
	e.GET("/user/all", GetAllUsers(uServcie))
	e.POST("/user/create", CreateUser(uServcie))
	e.POST("/user/login", LoginHandler(uServcie))
	return e
}

func CreateUser(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req user.UserRegisterDTO
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		// if err := c.Validate(&req); err != nil {
		// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		// }
		create, err := s.CreateUser(c.Request().Context(), req)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
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

func LoginHandler(s user.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req user.LoginUserDTO

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		// query to get user by email
		userPayload, err := s.GetUserByEmail(c.Request().Context(), req.Email)

		// verificando se o usuário é encontrado
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "user not found"})
		}

		// compare password hash
		if err := utils.VerifyHashPassword(req.Password, userPayload.Password); !err {
			return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Invalid Credentials"})
		}

		// Generate Token
		convertUUID, err := tools.ConvertUUIDToString(userPayload.UserID)
		if err != nil {
			return fmt.Errorf("erro to conevrt uuid to string: %v", err)
		}
		tokenString, err := utils.GenerateToken(convertUUID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create token"})
		}

		return c.JSON(http.StatusOK, tokenString)
	}
}
