package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/milinches/go-qrcode/controllers"
)

var fileServer = http.FileServer(http.Dir("./static"))

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/generate", controllers.QRGenerator).Methods("POST")
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))
	return router
}
