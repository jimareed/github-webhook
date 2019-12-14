package main

import (
	"net/http"
	"io"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	io.WriteString(w, "{\"id\": \"123456\"}\n")
}

