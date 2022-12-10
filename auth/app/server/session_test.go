package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewSessionManager(t *testing.T) {
	s := NewSessionManager("session_id")
	s.Set("key1", "value1")
	actual := s.Get("key1").(string)
	expected := "value1"
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestDeleteSession(t *testing.T) {
	s := NewSessionManager("session_id")
	s.Set("key1", "value1")
	s.Delete("key1")
	actual := s.Get("key1")
	expected := interface{}(nil)
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestSessionID(t *testing.T) {
	s := NewSessionManager("session_id")
	sid := s.SessionID()
	actual := len(sid)
	expected := 44
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}

func TestNewSession(t *testing.T) {
	expected := "session_id"
	s := NewSessionManager(expected)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.NewSession(w, r)
	}))
	defer srv.Close()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	srv.Config.Handler.ServeHTTP(w, r)
	cookies := w.Result().Cookies()
	cookie := cookies[0]
	actual := cookie.Name
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
	if !s.IsAlive(cookie.Value) {
		t.Errorf("session %v is not alive", cookie.Value)
	}
}
