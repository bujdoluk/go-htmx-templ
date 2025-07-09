package main

import (
	"expvar"
	"flag"
	"fmt"
	"os"

	"github.com/bujdoluk/go-htmx-templ/internal/logger"
)

var (
	version string
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *logger.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production)")

	displayVersion := flag.Bool("version", false, "Display version and exit")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		os.Exit(0)
	}

	expvar.NewString("version").Set(version)

	logger := logger.New(os.Stdout, logger.LevelInfo)

	app := &application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
