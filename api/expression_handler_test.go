package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func baseTest(expr string, expected string, t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(HandlePostExpression))
	defer server.Close()

	reqBody := strings.NewReader(expr)

	res, err := http.Post(server.URL, "text", reqBody)
	if err != nil {
		t.Error(err)
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	result := string(bytes)

	if expected != result {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSimpleExpressions(t *testing.T) {
	baseTest("4+5", "9", t)
	baseTest("5-4", "1", t)
	baseTest("5*4", "20", t)
	baseTest("20/4", "5", t)
}

func TestComplexExpressions(t *testing.T) {
	baseTest("(4-5)*(8+3)", "-11", t)
}

func TestErroneousExpressions(t *testing.T) {
	baseTest("())4*5", "Erroneous Expression - Remember using URL escape characters", t)
	baseTest("5:44 == 4", "Erroneous Expression - Remember using URL escape characters", t)
}
