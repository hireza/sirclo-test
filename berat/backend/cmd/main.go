package main

import (
	"fmt"
	"os"

	"github.com/hireza/sirclo-test/berat/packages"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/hireza/sirclo-test/berat/app/delivery/rest/route"
)

func run() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// START SERVICE
	mgr, err := packages.NewInit("", "")
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	server := mgr.GetServer()
	route.NewWeightRoute(server.Router, mgr)

	server.RegisterRouter(server.Router)
	return server.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
