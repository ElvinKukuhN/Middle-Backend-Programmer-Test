// routes/routes.go

package routes

import (
	"github.com/ElvinKukuhN/Middle-Backend-Programmer-Test/controller"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/convert/pngtojpeg", controller.ConvertPNGtoJPEG)
	http.HandleFunc("/resize", controller.ResizeImage)
	http.HandleFunc("/compress", controller.CompressImage)
}
