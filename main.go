package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/milinches/go-qrcode/routes"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env file")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("No env file found")
	}

	routes := routes.NewRouter()
	log.Println("Starting development server")
	err := http.ListenAndServe(":"+port, routes)
	log.Fatal(err)
}
