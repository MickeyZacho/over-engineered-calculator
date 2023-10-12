package main

import (
	"fmt"
	"net/http"
	"over-engineered-calculator/api"
	"over-engineered-calculator/views"
)

func main() {

	http.HandleFunc("/", views.HandleViewCalculator)
	http.HandleFunc("/expression", api.HandlePostExpression)
	http.HandleFunc("/log", api.HandleGetLog)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Something went wrong with creating the server")
		panic(err)
	}

}