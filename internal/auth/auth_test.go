package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Valid Authorization Header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey validapikey123")

		apiKey, err := GetAPIKey(headers)

		if err != nil || apiKey != "validapikey123" {
			t.Errorf("Expected apiKey: validapikey123, got: %s, error: %v", apiKey, err)
		}
	})
}
