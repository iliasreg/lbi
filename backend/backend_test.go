package backend

import (
	"testing"
)

func TestNewBackend(t *testing.T) {
	rawURL := "http://localhost:6969"
	backend := CreateBackend(rawURL, 100000)

	if backend.URL.String() != rawURL {
		t.Errorf("Expected URL %s, got %s", rawURL, backend.URL.String())
	}

	if !backend.IsAlive() {
		t.Error("Expected backend to be alive, but it's dead")
	}
}

func TestSetAlive(t *testing.T) {
	rawURL := "http://localhost:6969"
	backend := CreateBackend(rawURL, 10000)

	backend.SetAlive(false)
	if backend.IsAlive() {
		t.Error("Expected backend to be dead, but it's alive")
	}

	backend.SetAlive(true)
	if !backend.IsAlive() {
		t.Error("Expected backend to be alive, but it's dead")
	}
}
