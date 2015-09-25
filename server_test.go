package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCanvases(t *testing.T) {
	req, err := http.NewRequest("GET", "/canvases", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := &wallHandler{}
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("/canvases didn't return %v", http.StatusOK)
	}
}
