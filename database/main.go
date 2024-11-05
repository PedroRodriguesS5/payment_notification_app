package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "github.com/pedroRodriguesS5/payment_notification/database/db_config"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Erro ao capturar as variáveis de ambiente", err)
	}
	cfg := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Dbname:   os.Getenv("DB_NAME"),
		Sslmode:  "disable",
	}

	pool, err := db.ConnectPgx(cfg)

	if err != nil {
		log.Fatal("Erro ao estabelecer conexão com o banco de dados", err)
	}
	defer pool.Close()

	fmt.Println("Conexão bem sucedida")
}
