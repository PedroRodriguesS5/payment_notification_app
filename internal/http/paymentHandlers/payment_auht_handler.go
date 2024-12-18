package paymenthandlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
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

// @Summary Create payment
// @Description Create a recurring payment
// @Accept json
// @Produce json
// @Param data body payment.RecurringPaymentRequestDTO true "Post request body"
// @Success 201 {object} map[string]string "Pagamento criado com sucesso"
// @Failure 400 {object} map[string]string "error" : "Invalid request"
// @Failure 500 {object} map[string]string "error" : "Internal Server Error"
// @Router /public/user/create [post]
func CreatePayment(s payment.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("token").(string)
		if !ok || token == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "missing or expired token"})
		}
		var req payment.RecurringPaymentRequestDTO

		validationOk, err := govalidator.ValidateStruct(req)
		if !validationOk {
			return c.JSON(http.StatusBadRequest, map[string]error{"Invvalid Credetials": err})
		}

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		create, err := s.CreateRecurringPayments(c.Request().Context(), req, token)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]error{"error": err})
		}

		return c.JSON(http.StatusCreated, create)
	}
}

// @Summary Create self payment
// @Description Create a recurring self payment
// @Accept json
// @Produce json
// @Param data body payment.RecurringPaymentRequestDTO true "Post request body"
// @Success 201 {object} map[string]string "Pagamento criado com sucesso"
// @Failure 400 {object} map[string]string "error" : "Invalid request"
// @Failure 500 {object} map[string]string "error" : "Internal Server Error"
// @Router /public/user/create [post]
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

// @Summary Get payment user
// @Description Get a payment that user id is equal to payer id
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Produce json
// @Success 200 {array} payment.RecurringPaymentResponseDTO
// @Failure 400 {object} map[string]string "error" : "Payment not found"
// @Failure 500 {object} map[string]string "error" 	: "Internal server error"
// @Router /public/user/create [get]
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
