package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(file)
}

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to get uploaded file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	convertedStr, err := service.Detection(string(fileBytes))
	if err != nil {
		http.Error(w, "failed to convert file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	extension := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("%s_converted%s", time.Now().UTC().Format("20060102_150405"), extension)

	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "failed to create output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(convertedStr)
	if err != nil {
		http.Error(w, "failed to write to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(convertedStr))
}
