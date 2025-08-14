package auth

import (
	"testing"
	"net/http"
	"reflect"
	"errors"
)

func TestShouldSuccedd(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")
	got, err := GetAPIKey(headers)
	want := "abc123"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected key %q, got %q, error: %v", want, got, err)
	}
}

func TestNoKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey")
	got, err := GetAPIKey(headers)
	if err == nil{
		t.Fatalf("expected to fail but got %v", got)
	}
}

func TestMissingAuthHeader(t *testing.T) {
	headers := http.Header{}
	got, err := GetAPIKey(headers)
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	if got == "" {
		t.Fatalf("expected empty key, got %q", got)
	}
}