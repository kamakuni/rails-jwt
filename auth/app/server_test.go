package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

var s *http.Server

func TestMain(m *testing.M) {
	secret, _ := ReadSecret("../certs/private.key")
	s = NewAuthServer(":8080", secret)
	go func() {
		log.Fatal(s.ListenAndServe())
	}()
	c := m.Run()
	os.Exit(c)
}

func TestNewAuthServer(t *testing.T) {
	actual := s.Addr
	expected := ":8080"
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestAuth(t *testing.T) {
	user := &User{
		Email:    "test@example.com",
		Password: "password",
	}
	userJson, _ := json.Marshal(user)
	res, err := http.Post("http://localhost:8080/api/v1/auth", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "{\"token\":\"auth token\"}"
	if string(body) != expected {
		t.Errorf("response is not '%v'. actual:%v\n", expected, string(body))
	}
}

func TestRefresh(t *testing.T) {
	res, err := http.Get("http://localhost:8080/api/v1/refresh")
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	expected := "refresh token"
	if string(body) != expected {
		t.Errorf("response is not '%v'. actual:%v\n", expected, string(body))
	}
}
