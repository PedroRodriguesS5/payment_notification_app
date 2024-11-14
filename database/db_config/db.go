package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

func ConnStr(cfg Config) string {
	var strConn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname, cfg.Sslmode)

	return strConn
}

func ConnectPgx(cfg Config) (*pgxpool.Pool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := pgxpool.New(ctx, ConnStr(cfg))

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return conn, nil
}

func NewDbPgx(cfg Config) (*DB, error) {
	db, err := ConnectPgx(cfg)
	if err != nil {
		return nil, err
	}
	return &DB{Pool: db}, nil
}

type DB struct {
	Pool *pgxpool.Pool
}

func (d *DB) Ping() error {
	return d.Pool.Ping(context.Background())
}

func (d *DB) Close() error {
	d.Pool.Close()
	return nil
}
