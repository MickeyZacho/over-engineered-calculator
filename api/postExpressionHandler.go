package api

import (
	"fmt"
	"go/token"
	"go/types"
	"io"
	"net/http"
)

type User struct {
	Name string
	Id   int
}

type UserLog struct {
	User User
	Log  []LogEntry
}

type LogEntry struct {
	Expression string
	Result     string
}

var ExpressionLog []LogEntry

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

	ExpressionLog = append(ExpressionLog, LogEntry{Expression: expr, Result: result.Value.String()})

	io.WriteString(w, result.Value.String())

}
