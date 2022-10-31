package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ResponseToken struct {
	Token string `json:"token"`
}

func CreateToken(userId string, now time.Time, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "123456789",
		"exp": now.Add(10 * time.Minute).Unix(),
		"iat": now.Unix(),
	})
	return token.SignedString([]byte(secret))
}

func ReadSecret(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	data := make([]byte, 0124)
	count, err := f.Read(data)
	if err != nil {
		return "", err
	}
	return string(data[:count]), nil
}

func NewAuthServer(addr string, secret string) *http.Server {
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
		token, err := CreateToken("", time.Now(), secret)
		bytes, err := json.Marshal(&ResponseToken{
			Token: token,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(bytes))
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
