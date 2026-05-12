package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Benson-14/file-upload-service/internal/storage"
)

type Handler struct {
	S3 *storage.S3Client
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	slog.Info("The server is Healthy")
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, fileHandler, err := r.FormFile("file")
	if err != nil {
		slog.Error("error getting the file: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = h.S3.Upload(r.Context(), fileHandler.Filename, file, fileHandler.Size)
	if err != nil {
		slog.Error("error uploading file to S3: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "File uploaded successfully")
}
