package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func baseLogHandlerTest(expected []LogEntry, t *testing.T) {
	//expected := []LogEntry{{Expression: "4", Result: "4"}, {Expression: "5", Result: "5"}}

	ExpressionLog = []LogEntry{{Expression: "4", Result: "4"}, {Expression: "5", Result: "5"}}
	server := httptest.NewServer(http.HandlerFunc(HandleGetLog))
	defer server.Close()

	res, err := http.Get(server.URL + "/log")
	if err != nil {
		t.Error(err)
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	var result []LogEntry
	json.Unmarshal(bytes, &result)
	//fmt.Println(result)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

func TestLogHandler(t *testing.T) {
	x := []LogEntry{{Expression: "4", Result: "4"}, {Expression: "5", Result: "5"}}
	baseLogHandlerTest(x, t)
}
