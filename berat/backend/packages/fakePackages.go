package packages

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/hireza/sirclo-test/berat/packages/config"
	"github.com/hireza/sirclo-test/berat/packages/server"
)

type FakePackages interface {
	Packages
}

type fakePackages struct {
	config   *config.Config
	server   *server.Server
	postgres *sql.DB
}

func NewFakeInit(ctrl *gomock.Controller) (FakePackages, error) {
	cfg := &config.Config{}

	srv := server.NewServer(cfg)

	dbPostgres, _, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	return &fakePackages{
		config:   cfg,
		server:   srv,
		postgres: dbPostgres,
	}, nil
}

func (fp *fakePackages) GetConfig() *config.Config {
	return fp.config
}

func (fp *fakePackages) GetServer() *server.Server {
	return fp.server
}

func (fp *fakePackages) GetPostgres() *sql.DB {
	return fp.postgres
}
