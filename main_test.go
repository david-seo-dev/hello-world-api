package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HealthCheckHandler)    
    handler.Router.ServeHTTP(rr, req)

    return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func TestIndex(t *testing.T) {
    req, _ := http.NewRequest("GET", "/", nil)
    response := executeRequest(req)
    checkResponseCode(t, http.StatusOK, response.Code)
    
    if body := response.Body.String(); body != "[]" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}
