package paymenthandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedroRodriguesS5/payment_notification/internal/middleware"
	"github.com/pedroRodriguesS5/payment_notification/internal/service/payment"
)

func RegisterPaymentAuthRoutes(e *echo.Echo, pService payment.Service) {
	authGroup := e.Group("/auth")

	authGroup.Use(middleware.AuthToken)
	authGroup.POST("/payment/create", CreatePayment(pService))
	authGroup.POST("/payment/self/create", CreateSelfPayment(pService))
	authGroup.GET("/payment", GetPayment(pService))
}

func CreatePayment(s payment.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("token").(string)
		if !ok || token == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing or expired token"})
		}
		var req payment.RecurringPaymentRequestDTO
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		create, err := s.CreateRecurringPayments(c.Request().Context(), req, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error to create payment"})
		}

		return c.JSON(http.StatusCreated, create)
	}
}

func CreateSelfPayment(s payment.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("token").(string)
		if !ok || token == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing or expired token"})
		}
		var req payment.SelfPaymentDTO
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		create, err := s.CreateSelfRecurringPayment(c.Request().Context(), req, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, create)
	}
}

func GetPayment(s payment.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("token").(string)
		if !ok || token == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing or expired token"})
		}
		payment, err := s.GetRecurringPayement(c.Request().Context(), token)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Payment not found"})
		}

		return c.JSON(http.StatusOK, payment)
	}
}
