package main

import (
	"crypto/sha256"
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

	h := sha256.New()
	h.Write([]byte("123" + "very-secret-salt-that-should-live-in-.env"))
	passHash := h.Sum(nil)
	fmt.Println(passHash)

	fmt.Printf("eval: %g \n", res)

	http.HandleFunc("/", views.HandleViewCalculator)
	http.HandleFunc("/expression", api.HandlePostExpression)
	http.HandleFunc("/log", api.LogHandler)
	http.HandleFunc("/custom-eval", api.HandleCustomEval)
	http.HandleFunc("/login", api.HandleLogin)
	http.HandleFunc("/create-user", api.HandleCreateUser)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Something went wrong with creating the server")
		panic(err)
	}

}
