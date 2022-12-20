package server

import (
	"auth/constant"
	"auth/ent/authorizationcode"
	"auth/ent/enttest"
	"auth/ent/oauthclient"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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

func newHTTPClientWithoutRedirect() *http.Client {
	c := &http.Client{}
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return c
}

func createOAuthClient(clientName string, redirectURI string, scope string) (resp *http.Response, err error) {
	req := &RequestClient{
		ClientName:  clientName,
		RedirectURI: redirectURI,
		Scope:       scope,
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	res, err := http.Post("http://localhost:8080/api/v1/client", "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}
	return res, err
}

func TestClient(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	s.client = client

	res, err := createOAuthClient("javascript app", "https://localhost:3000/callback", "read write")
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

func TestGetAuthorize(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	s.client = client
	defer client.Close()

	res, _ := createOAuthClient("javascript app", "https://localhost:3000/callback", "read write")
	body, _ := io.ReadAll(res.Body)
	var resJSON ResponseClient
	if err := json.Unmarshal(body, &resJSON); err != nil {
		t.Error(err)
	}

	u, _ := url.Parse("http://localhost:8080/api/v1/authorize")
	params := u.Query()
	params.Add(constant.ResponseType.String(), constant.Code.String())
	params.Add(constant.Scope.String(), "read")
	params.Add(constant.State.String(), "abc")
	params.Add(constant.ClientID.String(), resJSON.ClientID)
	u.RawQuery = params.Encode()

	httpClient := newHTTPClientWithoutRedirect()
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	res, _ = httpClient.Do(req)
	l := res.Header.Get("Location")
	if !strings.HasPrefix(l, "https://localhost:3000/callback?code=") {
		t.Error("Unexpected redirect uri.")
	}
	if res.StatusCode != 302 {
		t.Error("Unexpected status code.")
	}
	ctx := context.Background()
	exist, _ := s.client.AuthorizationCode.Query().
		Where(authorizationcode.ClientID(resJSON.ClientID)).
		Exist(ctx)
	if !exist {
		t.Error("Authorization code is not found.")
	}
}

func TestPostAuthorize(t *testing.T) {

	v := url.Values{}
	v.Add("consent", "1")
	v.Add("redirect_uri", "http://localhost:8081/callback")
	httpClient := newHTTPClientWithoutRedirect()
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/authorize", strings.NewReader(v.Encode()))
	res, _ := httpClient.Do(req)
	l := res.Header.Get("Location")
	if !strings.HasPrefix(l, "http://localhost:8081/callback") {
		t.Error("Unexpected redirect uri.")
	}
	if res.StatusCode != 302 {
		t.Error("Unexpected status code.")
	}

}

func TestToken(t *testing.T) {
	res, err := http.Get("http://localhost:8080/api/v1/token")
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

func TestCreateTemplates(t *testing.T) {
	templates, _ := CreateTemplates("../template")
	tmpl, ok := templates["authorize.html"]
	if !ok {
		t.Error("template for authorize.html is not found.")
	}
	if tmpl.Name() != "authorize.html" {
		t.Error("template for authorize.html is not found.")
	}
	clientName := "javascript app"
	scopes := []string{"read", "write"}
	v := struct {
		ClientName string
		Scopes     []string
	}{
		ClientName: clientName,
		Scopes:     scopes,
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, v); err != nil {
		t.Error("Cannot set variables in template.")
	}
	fmt.Printf("%s", buf.String())
	if !(strings.Contains(buf.String(), clientName) &&
		strings.Contains(buf.String(), scopes[0]) &&
		strings.Contains(buf.String(), scopes[1])) {
		t.Error("Cannot set variables in template.")
	}
}
