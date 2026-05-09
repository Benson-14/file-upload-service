package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	srv := http.NewServeMux()

	slog.Info("Starting the server on port 8080...")
	err := http.ListenAndServe(":8080", srv)

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("server closed")
	} else if err != nil {
		slog.Error("error starting server", "err", err)
		os.Exit(1)
	}
}
