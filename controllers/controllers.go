package controllers

import (
	"html/template"
	"image/png"
	"log"
	"net/http"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	p := Page{Title: "QR CODE GENERATOR"}
	t, err := template.ParseFiles("views/qr.html")
	if err != nil {
		log.Println("HTML file not found.")
	}
	if err := t.Execute(w, p); err != nil {
		log.Fatal(err.Error())
	}
}

func QRGenerator(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("dataString")
	qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 300, 300)
	png.Encode(w, qrCode)
}
