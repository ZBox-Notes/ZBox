package database

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	DB *pgx.Conn
}

func NewPostgres() (*Postgres, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	if host == "" || user == "" || password == "" || database == "" {
		return nil, errors.New("missing required environment variables")
	}
	log.Println(host, user, password, database)
	db, err := pgx.Connect(context.Background(), "postgres://"+user+":"+password+"@"+host+"/"+database)
	if err != nil {
		return nil, err
	}
	return &Postgres{DB: db}, nil
}
