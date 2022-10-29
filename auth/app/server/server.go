package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func NewAuthServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body := make([]byte, length)
		length, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var jsonBody map[string]interface{}
		err = json.Unmarshal(body[:length], &jsonBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		email := jsonBody["email"]
		password := jsonBody["password"]
		if email != "test@example.com" || password != "password" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Printf("%v\n", jsonBody)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "auth token")
		return
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
