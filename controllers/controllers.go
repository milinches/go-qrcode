package controllers

import (
	"image/png"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func Home(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR CODE GENERATOR"}
	t, _ := template.ParseFiles("qr.html")
	t.Execute(w, p)
}

func QRGenerator(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("dataString")
	qrCode, _ := qr.Encode(data, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	png.Encode(w, qrCode)
}
