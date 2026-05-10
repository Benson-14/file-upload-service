package handler

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func Health(w http.ResponseWriter, r *http.Request) {
	slog.Info("The server is Healthy")
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		slog.Error("error getting the file: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// slog.Info("file name is " + fileHandler.Filename)
	// slog.Info(fmt.Sprintf("file size is %d bytes", fileHandler.Size))

	newFilePath := filepath.Join(os.Getenv("STORAGE"), fileHandler.Filename)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		slog.Error("error creating the file: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		slog.Error("error copying the file: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "File uploaded successfully")
}
