package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	db "github.com/pedroRodriguesS5/payment_notification/database/db_config"
	"github.com/pedroRodriguesS5/payment_notification/internal/handler/api"
	paymenthandlers "github.com/pedroRodriguesS5/payment_notification/internal/http/paymentHandlers"
	"github.com/pedroRodriguesS5/payment_notification/internal/http/userHandlers"
	"github.com/pedroRodriguesS5/payment_notification/internal/service/payment"
	"github.com/pedroRodriguesS5/payment_notification/internal/service/user"
	sqlc_db "github.com/pedroRodriguesS5/payment_notification/project"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/pedroRodriguesS5/payment_notification/docs"
)

// @title Payment Notification App
// @version 1.0
// @description notification from pending payments
// @contact.email pedroxbrs@gmail.com
// @host localhost:8000
// @basePath /public/user/login
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Erro ao capturar as variáveis de ambiente", err)
	}

	cfg := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: "Macaco584@@",
		Dbname:   os.Getenv("DB_NAME"),
		Sslmode:  "disable",
	}

	pool, err := db.ConnectPgx(cfg)

	if err != nil {
		log.Fatal("Erro ao estabelecer conexão com o banco de dados", err)
	}

	defer pool.Close()

	err = pool.Ping(context.Background())

	if err != nil {
		log.Fatalf("error to verify the connection: %v", err)
	}

	fmt.Println("Conection successful", pool.Stat())

	// Services
	queries := sqlc_db.New(pool)
	pService := payment.NewService(queries)
	uService := user.NewService(queries)
	// Echo instance
	e := echo.New()

	// h = echo.Handlres(*uService)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	paymenthandlers.RegisterPaymentAuthRoutes(e, *pService)
	userHandlers.RegisterUserPublicRoutes(e, *uService)
	userHandlers.RegisterUserAuthRoutes(e, *uService)
	err = api.Start("8000", e)

	if err != nil {
		log.Fatal("Error to runing api", err)
	}
}
