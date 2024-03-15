package controller

import (
	"bytes"
	"image/jpeg"
	"net/http"
	"strconv"
)

func CompressImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	qualityStr := r.FormValue("quality")
	quality, err := strconv.Atoi(qualityStr)
	if err != nil || quality < 1 || quality > 100 {
		http.Error(w, "Invalid quality parameter", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	// Kompresi gambar menggunakan buffer sebagai output
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	if err != nil {
		http.Error(w, "Failed to compress image", http.StatusInternalServerError)
		return
	}

	// Menulis kembali buffer ke ResponseWriter
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	if _, err := w.Write(buf.Bytes()); err != nil {
		http.Error(w, "Failed to write image to response", http.StatusInternalServerError)
		return
	}
}
