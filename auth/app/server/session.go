package server

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type SessionManager interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

type Session struct {
	cookieName string
	sessionMap map[interface{}]interface{}
}

func NewSession(cookieName string) *Session {
	return &Session{
		cookieName: cookieName,
		sessionMap: map[interface{}]interface{}{},
	}
}

func (s *Session) Set(key, value interface{}) error {
	if _, exists := s.sessionMap[key]; exists {
		return errors.New("session already exists.\n")
	}
	s.sessionMap[key] = value
	return nil
}

func (s *Session) Get(key interface{}) interface{} {
	return s.sessionMap[key]
}

func (s *Session) Delete(key interface{}) error {
	if _, exists := s.sessionMap[key]; !exists {
		return errors.New("session does not exist.\n")
	}
	delete(s.sessionMap, key)
	return nil
}

func (s *Session) SessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (s *Session) NewSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(s.cookieName)
	if err != nil || cookie.Value == "" {
		sid := s.SessionID()
		cookie := http.Cookie{Name: s.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true}
		http.SetCookie(w, &cookie)
	}
}
