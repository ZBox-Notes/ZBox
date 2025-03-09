package database

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
)

type Postgres struct {
	DB *pgx.Conn
}

func NewPostgres() (*Postgres, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	if user == "" || password == "" || database == "" {
		return nil, errors.New("missing required environment variables")
	}
	db, err := pgx.Connect(context.Background(), "postgres://"+user+":"+password+"@postgres/"+database)
	if err != nil {
		return nil, err
	}
	return &Postgres{DB: db}, nil
}

func NewMockPostgres() (*Postgres, error) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		return nil, err
	}
	return &Postgres{DB: mock.Conn()}, nil
}
