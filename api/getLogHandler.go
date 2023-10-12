package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandleGetLog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/log" {
		http.Error(w, "500 server error", http.StatusInternalServerError)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	jsonExpressionLog, _ := json.Marshal(ExpressionLog)

	io.WriteString(w, string(jsonExpressionLog))

}
