package client

import (
	"fmt"
	"go-news-clean/pkg/env"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() (*sqlx.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		env.Host,
		env.DbPort,
		env.User,
		env.Dbname,
		env.Password,
		env.SslMode,
	)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
