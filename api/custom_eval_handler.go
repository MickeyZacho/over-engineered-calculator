package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func HandleCustomEval(w http.ResponseWriter, r *http.Request) {
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

	result, err := Eval(expr)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	resultString := fmt.Sprintf("%f", result)
	ExpressionLog = append(ExpressionLog, LogEntry{Expression: expr, Result: resultString})
	io.WriteString(w, resultString)
}

func Eval(expr string) (float64, error) {
	errEval := errors.New("something went wrong evaluating an expression")
	splitExpr := strings.Split(expr, " ")

	// if length = 1 we should necessarily be dealing with a number (hopefully)
	if len(splitExpr) == 1 {
		num, err := strconv.ParseFloat(expr, 64)
		if err != nil {
			return 0, errEval
		}
		return num, nil
	}

	// handle higher precedence operators
	for i, x := range splitExpr {
		//Excessive unneeded string joining. should only happen when needed.
		//Excessive copy-pasting of a = eval(e1), b = eval(e2)
		expr1 := strings.Join(splitExpr[:i], " ")
		expr2 := strings.Join(splitExpr[i+1:], " ")
		switch x {
		case "*":
			a, err := Eval(expr1)
			if err != nil {
				return 0, err
			}
			b, err := Eval(expr2)
			if err != nil {
				return 0, err
			}
			return a * b, nil
		case "/":
			a, err := Eval(expr1)
			if err != nil {
				return 0, err
			}
			b, err := Eval(expr2)
			if err != nil {
				return 0, err
			}
			return a / b, nil
		}
	}

	// handle lower precedence operators
	for i, x := range splitExpr {
		//Excessive unneeded string joining. should only happen when needed.
		expr1 := strings.Join(splitExpr[:i], " ")
		expr2 := strings.Join(splitExpr[i+1:], " ")
		switch x {
		case "+":
			a, err := Eval(expr1)
			if err != nil {
				return 0, err
			}
			b, err := Eval(expr2)
			if err != nil {
				return 0, err
			}
			return a + b, nil
		case "-":
			a, err := Eval(expr1)
			if err != nil {
				return 0, err
			}
			b, err := Eval(expr2)
			if err != nil {
				return 0, err
			}
			return a - b, nil
		}
	}

	// If we have not already returned, something went wrong.
	fmt.Println("Something went wrong evaluating an expression")
	return 0, errEval

}
