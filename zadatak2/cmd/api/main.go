package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5000, "API server port")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	logger.Printf("starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)
}
