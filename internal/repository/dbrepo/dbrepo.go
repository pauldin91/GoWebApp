package dbrepo

import (
	"database/sql"

	"github.com/pauldin91/GoWebApp/internal/cfg"
	"github.com/pauldin91/GoWebApp/internal/repository"
)

type postgresDbRepo struct {
	App *cfg.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *cfg.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		App:a,
		DB:conn,
	}
}
