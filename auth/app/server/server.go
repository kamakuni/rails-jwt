package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"auth/constant"
	"auth/ent"
	"auth/ent/oauthclient"
)

type Server struct {
	*http.Server
	client *ent.Client
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

func NewAuthServer(ctx context.Context, client *ent.Client, addr string, secret string) *Server {
	s := &Server{
		Server: &http.Server{
			Addr: addr,
		},
		client: client,
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
		return
	})
	mux.HandleFunc("/api/v1/authorize", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		params := r.URL.Query()
		responseType := params.Get("response_type")
		if responseType != constant.Code.String() {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		clientID := params.Get("client_id")
		state := params.Get("state")
		//scope := params.Get("scope")
		c, err := s.client.OAuthClient.Query().
			Where(oauthclient.ClientID(clientID)).
			Only(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		redirectURI, err := url.Parse(c.RedirectURI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		values := redirectURI.Query()
		code, err := CreateCode()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.client.AuthorizationCode.Create().
			SetClientID(clientID).
			SetCode(code).
			Save(ctx)
		values.Add(constant.Code.String(), code)
		values.Add("state", state)
		redirectURI.RawQuery = values.Encode()
		//accessToken, err := CreateAccessToken("", time.Now(), secret)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//}
		//refreshToken, err := CreateRefreshToken()
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//bytes, err := json.Marshal(&ResponseAuthorize{
		//	AccessToken:  accessToken,
		//	RefreshToken: refreshToken,
		//})
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		w.Header().Add("Location", redirectURI.String())
		w.WriteHeader(http.StatusFound)
		return
	})
	mux.HandleFunc("/api/v1/refresh", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "refresh token")
	})
	s.Handler = mux
	return s
}
