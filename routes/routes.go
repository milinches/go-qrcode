package routes

import (
	"github.com/gorilla/mux"
	"github.com/milinches/go-qrcode/controllers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/generate", controllers.QRGenerator).Methods("POST")
	return router
}
