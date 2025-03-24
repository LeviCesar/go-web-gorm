package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-web/models"
)

var templateHtml = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Print("Running: http://localhost:7000")

	http.HandleFunc("/", index)
	http.ListenAndServe(":7000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	productList := product.Read()
	templateHtml.ExecuteTemplate(w, "Index", productList)
}
