package main

import (
	"fmt"
	"net/http"
	"over-engineered-calculator/api"
	"over-engineered-calculator/views"
)

func main() {

	res, err := api.Eval("4 + 5 * 6")
	if err != nil {
		fmt.Println("error happened")
	}

	fmt.Printf("eval: %g \n", res)

	http.HandleFunc("/", views.HandleViewCalculator)
	http.HandleFunc("/expression", api.HandlePostExpression)
	http.HandleFunc("/log", api.HandleGetLog)
	http.HandleFunc("/custom-eval", api.HandleCustomEval)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Something went wrong with creating the server")
		panic(err)
	}

}
