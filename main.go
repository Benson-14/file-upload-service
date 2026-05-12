package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Benson-14/file-upload-service/internal/config"
	"github.com/Benson-14/file-upload-service/internal/db"
	"github.com/Benson-14/file-upload-service/internal/handler"
	"github.com/Benson-14/file-upload-service/internal/middleware"
	"github.com/Benson-14/file-upload-service/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file: " + err.Error())
		os.Exit(1)
	}

	cfg := config.LoadConfig()

	_, err = db.NewDatabase(cfg.DB.URL)
	if err != nil {
		slog.Error("error connecting to database: " + err.Error())
		os.Exit(1)
	}

	s3Client, err := storage.NewS3Client(cfg)
	if err != nil {
		slog.Error("error initialising S3 client: " + err.Error())
		os.Exit(1)
	}

	h := &handler.Handler{S3: s3Client}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("POST /upload", h.UploadFile)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      middleware.RequestLogger(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	slog.Info(fmt.Sprintf("Starting the server on port %d...", cfg.App.Port))
	err = srv.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("server closed")
	} else if err != nil {
		slog.Error("error starting server: " + err.Error())
		os.Exit(1)
	}
}
