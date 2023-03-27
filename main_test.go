package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCmdHandler(t *testing.T) {
	// Create a test request with a valid command.
	cmd := Command{Cmd: "echo hello"}
	body, _ := json.Marshal(cmd)
	req := httptest.NewRequest(http.MethodPost, "/api/cmd", bytes.NewBuffer(body))

	// Create a test response recorder.
	w := httptest.NewRecorder()

	// Call the handler function with the test request and response recorder.
	cmdHandler(w, req)

	// Check that the response status code is OK (200) and the response body is "hello\n".
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
	if w.Body.String() != "hello\n" {
		t.Errorf("Expected body %q, but got %q", "hello\n", w.Body.String())
	}

	// Create a test request with an empty command.
	cmd = Command{Cmd: ""}
	body, _ = json.Marshal(cmd)
	req = httptest.NewRequest(http.MethodPost, "/api/cmd", bytes.NewBuffer(body))

	// Call the handler function with the test request and response recorder.
	w = httptest.NewRecorder()
	cmdHandler(w, req)

	// Check that the response status code is Not Found (404).
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
	}
}
