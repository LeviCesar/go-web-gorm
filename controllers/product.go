package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-web/models"
)

var templateHtml = template.Must(template.ParseGlob("templates/products/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	productList := product.List()
	templateHtml.ExecuteTemplate(w, "Index", productList)
}

func New(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templateHtml.ExecuteTemplate(w, "New", nil)
	} else if r.Method == "POST" {
		name := r.FormValue("name")
		price := r.FormValue("price")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price Convert Error: ", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity Convert Error: ", err)
		}

		product := models.Product{Name: name, Description: description, Price: priceConverted, Quantity: uint16(quantityConverted)}
		if product.Create() {
			fmt.Println("Product add success!")
		} else {
			fmt.Println("Product add fail!")
		}

		http.Redirect(w, r, "/product", http.StatusMovedPermanently)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	productIdConverted, err := strconv.Atoi(productId)
	if err != nil {
		log.Println("Product id not defined", productIdConverted)
	}

	product := models.Get(uint(productIdConverted))
	product.Delete()

	http.Redirect(w, r, "/product", http.StatusMovedPermanently)
}
