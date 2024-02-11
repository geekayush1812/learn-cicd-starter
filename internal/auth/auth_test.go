package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if key != "test-api-key" {
		t.Errorf("Incorrect API key extracted: expected 'test-api-key', got '%s'", key)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := http.Header{}

	key, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error '%v', got '%v'", ErrNoAuthHeaderIncluded, err)
	}

	if key != "" {
		t.Errorf("Incorrect API key extracted: expected '', got '%s'", key)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer some-token")

	key, err := GetAPIKey(headers)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if key != "" {
		t.Errorf("Incorrect API key extracted: expected '', got '%s'", key)
	}
}
