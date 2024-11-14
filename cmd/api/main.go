package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "github.com/pedroRodriguesS5/payment_notification/database/db_config"
	"github.com/pedroRodriguesS5/payment_notification/internal/handler/api"
	"github.com/pedroRodriguesS5/payment_notification/internal/http/echo"
	"github.com/pedroRodriguesS5/payment_notification/internal/user"
	sqlc_db "github.com/pedroRodriguesS5/payment_notification/project"
)

// conection to database and start the api
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

	queries := sqlc_db.New(pool)
	uService := user.NewService(queries)
	h := echo.Handlres(*uService)

	err = api.Start("8000", h)

	if err != nil {
		log.Fatal("Error to runing api", err)
	}
}