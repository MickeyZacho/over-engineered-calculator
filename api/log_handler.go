package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LogEntry struct {
	Expression string
	Result     string
}

var ExpressionLog []LogEntry

func LogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/log" {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {
		HandleGetLog(w, r)
	}
	if r.Method == http.MethodPost {
		HandlePostLog(w, r)
	}

}

func HandleGetLog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/log" {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	jsonExpressionLog, err := json.Marshal(ExpressionLog)
	if err != nil {
		io.WriteString(w, "error occured: "+err.Error())
	}

	io.WriteString(w, string(jsonExpressionLog))

}

func HandlePostLog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/log" {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	exprbytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Could not read request body")
	}
	var log LogEntry

	json.Unmarshal(exprbytes, &log)

	ExpressionLog = append(ExpressionLog, log)

	io.WriteString(w, "log recieved")

}
