package client

import (
	"fmt"
	"go-news-clean/pkg/env"
)

// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
var ConnString string = fmt.Sprintf(
	"postgres://%s:%s@%s:%s/%s",
	env.User,
	env.Password,
	env.Host,
	env.DbPort,
	env.Dbname,
)
