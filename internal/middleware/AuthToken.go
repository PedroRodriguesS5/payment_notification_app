package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/pkg/infra"
)

// Authorization token
func AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or expired token")
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := infra.VerifyToken(tokenString)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		// Attach claims to the context for later use
		c.Set("token", tokenString)
		c.Set("userClaims", claims)
		return next(c)
	}
}
