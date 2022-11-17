package server

import (
	"auth/ent/enttest"
	"auth/ent/oauthclient"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var s *Server

func TestMain(m *testing.M) {
	//client := Open("postgres://postgres:password@auth-db/postgres?sslmode=disable")
	//client := enttest.Open(m., "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	//defer client.Close()s
	ctx := context.Background()
	secret, _ := ReadSecret("../certs/private.key")
	s = NewAuthServer(ctx, nil, ":8080", secret)
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

func TestClient(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	s.client = client
	req := &RequestClient{
		ClientName:  "javascript app",
		RedirectURI: "https://localhost:3000/callback",
		Scope:       "read write",
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
	}
	res, err := http.Post("http://localhost:8080/api/v1/client", "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	var resJSON ResponseClient
	if err := json.Unmarshal(body, &resJSON); err != nil {
		t.Error(err)
	}
	if resJSON.ClientID == "" {
		t.Errorf("response has no ClientID.")
	}
	if resJSON.ClientName == "" {
		t.Errorf("response has no ClientName.")
	}
	ctx := context.Background()
	exist, _ := s.client.OAuthClient.
		Query().
		Where(oauthclient.ClientNameEQ("javascript app")).Exist(ctx)
	if !exist {
		t.Errorf("OAuth client is not found.")
	}
}

func TestToken(t *testing.T) {
	user := &User{
		Email:    "test@example.com",
		Password: "password",
	}
	userJson, _ := json.Marshal(user)
	res, err := http.Post("http://localhost:8080/api/v1/token", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	var responseJson Response
	if err := json.Unmarshal(body, &responseJson); err != nil {
		t.Error(err)
	}
	if responseJson.AccessToken == "" {
		t.Errorf("response has no AccessToken.")
	}
	if responseJson.RefreshToken == "" {
		t.Errorf("response has no RefreshToken.")
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
