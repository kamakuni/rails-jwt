package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/auth", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "auth token")
	})
	http.HandleFunc("/api/v1/refresh", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "refresh token")
	})
}
