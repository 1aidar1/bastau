package main

import (
	"1aidar1/bastau/go-api/config"
	"1aidar1/bastau/go-api/internal/app"
	"1aidar1/bastau/go-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	l := logger.NewLogger(logrus.New())

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg, l)
}
