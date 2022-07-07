package packages

import (
	"database/sql"

	"github.com/hireza/sirclo-test/berat/packages/config"
	"github.com/hireza/sirclo-test/berat/packages/database/postgres"
	"github.com/hireza/sirclo-test/berat/packages/server"
	"github.com/rs/zerolog/log"
)

type Packages interface {
	GetConfig() *config.Config
	GetServer() *server.Server
	GetPostgres() *sql.DB
}

type packages struct {
	config   *config.Config
	server   *server.Server
	postgres *sql.DB
}

func NewInit(path, host string) (Packages, error) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		log.Error().Err(err).Msg("error when initiate config")
		return nil, err
	}

	srv := server.NewServer(cfg)

	dbPostgres, err := postgres.NewPostgres(cfg).Connect(host)
	if err != nil {
		log.Error().Err(err).Msg("error when initiate database")
		return nil, err
	}

	return &packages{
		config:   cfg,
		server:   srv,
		postgres: dbPostgres,
	}, nil
}

func (p *packages) GetConfig() *config.Config {
	return p.config
}

func (p *packages) GetServer() *server.Server {
	return p.server
}

func (p *packages) GetPostgres() *sql.DB {
	return p.postgres
}
