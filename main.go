package main

import (
	"fmt"
	"net/http"

	"github.com/go-web/controllers"
)

func main() {
	fmt.Println("Running: http://localhost:7000")

	// Routes
	http.HandleFunc("/product", controllers.Index)
	http.HandleFunc("/product/new", controllers.New)
	http.HandleFunc("/product/delete", controllers.Delete)
	http.HandleFunc("/product/edit", controllers.Edit)
	// End Routes

	http.ListenAndServe(":7000", nil)
}
