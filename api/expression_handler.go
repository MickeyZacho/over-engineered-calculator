package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/token"
	"go/types"
	"io"
	"net/http"
)

type UserLog struct {
	User User
	Log  []LogEntry
}

func HandlePostExpression(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	exprbytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Could not read request body")
	}
	expr := string(exprbytes)

	fmt.Println("Received expr: " + expr)

	fs := token.NewFileSet()
	result, err := types.Eval(fs, nil, token.NoPos, expr)
	if err != nil {
		io.WriteString(w, "Erroneous Expression - Remember using URL escape characters")
		return
	}

	// Creating log and making it a *bytes.Reader
	log := LogEntry{Expression: expr, Result: result.Value.String()}
	logbytes, err := json.Marshal(log)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	bodyReader := bytes.NewReader(logbytes)

	// sending log to the logservice
	url := "http://localhost:8080/log"
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, err.Error())
		return
	}
	client.Do(req)

	io.WriteString(w, result.Value.String())

}
