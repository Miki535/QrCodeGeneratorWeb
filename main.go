package main

import (
	"html/template"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
}

func GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// parsing url from html
		url := r.FormValue("url")
		// error if url nil
		if url == "" {
			http.Error(w, "Please enter information", http.StatusBadRequest)
			return
		}
		// generate qr code
		qr, err := qrcode.Encode(url, qrcode.Medium, 256)
		// error while generating qrcode
		if err != nil {
			http.Error(w, "Error generating QR code", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(qr)
	}
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/generate", GenerateQRCode)

	http.ListenAndServe(":8080", nil)
}
