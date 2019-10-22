package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestIndex(t *testing.T) {
    req, _ := http.NewRequest("GET", "/", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    
    if body := response.Body.String(); body != "[]" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}
