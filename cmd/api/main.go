package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bujdoluk/go-htmx-templ/internal/logger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

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
