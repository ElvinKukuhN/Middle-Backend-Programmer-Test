package controller

import (
	"github.com/disintegration/imaging"
	"image"
	"net/http"
	"os"
)

func ConvertPNGtoJPEG(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	jpegImage, err := os.CreateTemp("", "*.jpeg")
	if err != nil {
		http.Error(w, "Failed to create temporary file", http.StatusInternalServerError)
		return
	}
	defer os.Remove(jpegImage.Name())

	err = imaging.Encode(jpegImage, img, imaging.JPEG)
	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, jpegImage.Name())
}
