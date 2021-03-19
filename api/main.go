package main

import (
	"github.com/nkarpenko/playlog-test/api/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		log.WithError(err).Fatal("Unexpected error in the service.")
	}
}
