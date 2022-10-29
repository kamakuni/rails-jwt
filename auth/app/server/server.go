package server

import (
	"io"
	"net/http"
)

func NewAuthServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/auth", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "auth token")
	})
	mux.HandleFunc("/api/v1/refresh", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "refresh token")
	})
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return s
}
