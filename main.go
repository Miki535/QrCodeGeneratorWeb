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
		url := r.FormValue("url")
		if url == "" {
			http.Error(w, "Please provide a URL", http.StatusBadRequest)
			return
		}

		qr, err := qrcode.Encode(url, qrcode.Medium, 256)
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
