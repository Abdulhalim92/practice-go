package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPServer(t *testing.T) {
	const numRequest = 30

	// Запуск HTTP - сервера в тестовом режиме
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
}
