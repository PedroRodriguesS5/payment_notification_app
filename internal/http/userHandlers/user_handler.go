package userHandlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/internal/service/user"
	"github.com/pedroRodriguesS5/payment_notification/pkg/infra"
	tools "github.com/pedroRodriguesS5/payment_notification/pkg/utils"
)

// Handler who the echo will create the routes
func RegisterUserPublicRoutes(e *echo.Echo, uService user.Service) {
	publicGroup := e.Group("/public")
	publicGroup.POST("/user/create", CreateUser(uService))
	publicGroup.POST("/user/login", LoginHandler(uService))
}

// @Summary Create User
// @Description Create a user in the database
// @Accept json
// @Produce json
// @Param data body user.UserRegisterDTO true "Post request body"
// @Success 201 {object} map[string]string "Usuário criado com sucesso"
// @Failure 400 {object} map[string]string "error" : "Invalid Request"
// @Failure 500 {object} map[string]string "error" : "Internal Server Error"
// @Router /public/user/create [post]
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

// @Summary User login
// @Description login and generate token for the user
// @Accept json
// @Param data body user.LoginUserDTO true "Login request body"
// @Success 200 {string} string "Return token"
// @Failure 400 {object} map[string]string "error" : "Invalid Credentails"
// @Failure 500 {object} map[string]string "error" : "Token generation error
// @Router /public/user/login [post]
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
			return c.JSON(http.StatusBadRequest, map[string]string{"erro": "Invalid Credentials"})
		}

		// compare password hash
		if err := infra.VerifyHashPassword(req.Password, userPayload.Password); !err {
			return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Invalid Credentials"})
		}

		// Generate Token
		convertUUID, err := tools.ConvertUUIDToString(userPayload.UserID)
		if err != nil {
			return fmt.Errorf("erro to conevrt uuid to string: %v", err)
		}
		tokenString, err := infra.GenerateToken(convertUUID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Token generate error"})
		}

		return c.JSON(http.StatusOK, tokenString)
	}
}
