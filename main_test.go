package main

import (
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    // Simulate an HTTP request to "/"
    req := httptest.NewRequest("GET", "/", nil)
   

    // Record the response using httptest
    w := httptest.NewRecorder()

    // Call the hander function
    handler(w, req)

    // Expected response
    expected := "Hello, CI/CD from Go!"
    actual := w.Body.String()

    if actual != expected {
        t.Errorf("Expected '%s', but go '%s'", expected, actual)
    }
}

