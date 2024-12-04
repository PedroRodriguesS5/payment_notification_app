package paymenthandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/internal/middleware"
	"github.com/pedroRodriguesS5/payment_notification/internal/payment"
)

func RegisterPaymentAuthRoutes(e *echo.Echo, pService payment.Service) {
	authGtoup := e.Group("/auth")

	authGtoup.Use(middleware.AuthToken)
	authGtoup.POST("/payment/create", CreatePayment(pService))
}

func CreatePayment(s payment.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("token").(string)
		if !ok || token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing or expired token"})
		}
		var req payment.RecurringPaymentRequestDTO
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		create, err := s.CreateRecurringPayments(c.Request().Context(), req, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, create)
	}
}
