package views

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func HandleViewCalculator(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "calculator", Body: []byte("This is a sample Page.")}
	t, err := template.ParseFiles("views/calculator.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, p)
}
