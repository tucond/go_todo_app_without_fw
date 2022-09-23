package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tucond/go_todo_app_without_fw/config"
	"honnef.co/go/tools/config"
)

func New(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword,
			cfg.DBHost, cfg.DBPort,
			cfg.DBName
		),
	)

	if err!=nil{
		return nil, nil, err
	}
}

