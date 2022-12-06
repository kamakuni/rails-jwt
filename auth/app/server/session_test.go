package server

import "testing"

func TestNewSession(t *testing.T) {
	s := NewSession("session_id")
	s.Set("key1", "value1")
	actual := s.Get("key1").(string)
	expected := "value1"
	if actual != expected {
		t.Errorf("actual:%v, expected:%v", actual, expected)
	}
}
