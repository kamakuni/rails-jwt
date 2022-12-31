package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"auth/constant"
	"auth/ent"
	"auth/ent/authorizationcode"
	"auth/ent/oauthclient"
	"auth/ent/refreshtoken"
)

type Server struct {
	*http.Server
	client    *ent.Client
	templates map[string]*template.Template
	session   *Session
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

func CreateTemplates(tmpldir string) (map[string]*template.Template, error) {
	templates := map[string]*template.Template{}
	files, err := filepath.Glob(filepath.Join(tmpldir, "*.html"))
	if err != nil {
		return templates, err
	}
	for _, file := range files {
		name := filepath.Base(file)
		tmpl := template.Must(template.New(name).ParseFiles(file))
		templates[name] = tmpl
	}
	return templates, nil
}

func NewAuthServer(ctx context.Context, client *ent.Client, addr string, secret string) *Server {
	session := NewSessionManager("session_id")
	templates, _ := CreateTemplates("../template")
	s := &Server{
		Server: &http.Server{
			Addr: addr,
		},
		client:    client,
		templates: templates,
		session:   session,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/client", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body := make([]byte, length)
		length, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var jsonBody map[string]interface{}
		err = json.Unmarshal(body[:length], &jsonBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientID, err := CreateClientID()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientName := jsonBody["client_name"].(string)
		redirectURI := jsonBody["redirect_uri"].(string)
		scope := jsonBody["scope"].(string)
		if clientName == "" || redirectURI == "" {
			http.Error(w, "", http.StatusUnprocessableEntity)
			return
		}
		c, err := s.client.OAuthClient.
			Create().
			SetClientID(clientID).
			SetClientName(clientName).
			SetClientType(constant.Public.String()).
			SetRedirectURI(redirectURI).
			SetScope(scope).
			Save(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("client: %v", c)
		res := &ResponseClient{ClientID: clientID, ClientName: clientName}
		buf, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	})
	mux.HandleFunc("/api/v1/authorize", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			params := r.URL.Query()
			responseType := params.Get("response_type")
			if responseType != constant.Code.String() {
				http.Error(w, "", http.StatusInternalServerError)
				return
			}
			clientID := params.Get("client_id")
			state := params.Get("state")
			codeChallenge := params.Get("code_challenge")
			scope := params.Get("scope")
			c, err := s.client.OAuthClient.Query().
				Where(oauthclient.ClientID(clientID)).
				Only(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			code, err := CreateCode()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			s.client.AuthorizationCode.Create().
				SetClientID(clientID).
				SetCode(code).
				SetCodeChallenge(codeChallenge).
				SetCodeChallengeMethod("plain").
				Save(ctx)
			redirectURI, err := url.Parse(c.RedirectURI)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			values := redirectURI.Query()
			values.Add(constant.Code.String(), code)
			values.Add("state", state)
			values.Add("scope", scope)
			redirectURI.RawQuery = values.Encode()
			tmpl, ok := s.templates["authorize.html"]
			if !ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data := struct {
				ClientName  string
				Scopes      []string
				RedirectURI string
			}{
				ClientName:  c.ClientName,
				Scopes:      strings.Split(scope, " "),
				RedirectURI: redirectURI.String(),
			}
			tmpl.Execute(w, data)
			return
		} else if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			consent := r.Form.Get("consent")
			if consent != "1" {
				http.Error(w, "", http.StatusBadRequest)
				return
			}
			redirectURI := r.Form.Get("redirect_uri")
			w.Header().Add("Location", redirectURI) //redirectURI.String())
			w.WriteHeader(http.StatusFound)
			return
		} else {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

	})
	mux.HandleFunc("/api/v1/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		length, err := strconv.Atoi(r.Header.Get("Content-Length"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body := make([]byte, length)
		length, err = r.Body.Read(body)
		if err != nil && err != io.EOF {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var jsonBody map[string]interface{}
		err = json.Unmarshal(body[:length], &jsonBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientID := jsonBody["client_id"].(string)
		code := jsonBody["code"].(string)
		refreshToken := jsonBody["refresh_token"].(string)
		if refreshToken != "" {
			_, err := s.client.RefreshToken.Query().
				Where(refreshtoken.Token(refreshToken)).
				Only(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// TODO valid refreshToken
			accessToken, err := CreateAccessToken("", time.Now(), secret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			buf, err := json.Marshal(&ResponseAuthorize{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(buf)
			return
		} else {
			_, err = s.client.AuthorizationCode.Query().
				Where(authorizationcode.ClientID(clientID), authorizationcode.Code(code)).
				Only(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			accessToken, err := CreateAccessToken("", time.Now(), secret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			refreshToken, err := CreateRefreshToken()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			buf, err := json.Marshal(&ResponseAuthorize{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(buf)
			return
		}
	})
	s.Handler = mux
	return s
}
