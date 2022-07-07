package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/hireza/sirclo-test/berat/packages/config"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

type Postgres interface {
	Connect(host string) (*sql.DB, error)
}

type Options struct {
	dbname string
	dbuser string
	dbpass string
}

func NewPostgres(cfg *config.Config) Postgres {
	opt := new(Options)
	opt.dbname = cfg.Database.DBName
	opt.dbuser = cfg.Database.DBUser
	opt.dbpass = cfg.Database.DBPass

	return opt
}

func (o *Options) Connect(host string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@db/%s?sslmode=disable", o.dbuser, o.dbpass, o.dbname)
	if host != "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", o.dbuser, o.dbpass, host, o.dbname)
	}

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("error when connect database")
		return nil, err
	}

	database.SetMaxOpenConns(50)
	database.SetMaxIdleConns(20)
	database.SetConnMaxLifetime(time.Hour)
	return database, nil
}
