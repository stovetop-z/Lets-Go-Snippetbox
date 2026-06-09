package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Get the address from the command line
	var config Config
	config.ConfigureFlags()

	// Initialize logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	app := application{
		logger: logger,
	}

	// Initialize the http function handler object
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(config.staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.getSnippetView)
	mux.HandleFunc("GET /snippet/create", app.getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", app.postSnippetCreate)

	logger.Info("starting server", "HTTP network address", config.addr)
	err := http.ListenAndServe(config.addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
