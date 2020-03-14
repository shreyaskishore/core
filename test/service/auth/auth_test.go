package auth_test

import (
	"testing"
	"time"

	"github.com/acm-uiuc/core/service/auth"
)

const (
	provider = "fake"
)

func TestGetOAuthRedirect(t *testing.T) {
	svc, err := auth.New()
	if err != nil {
		t.Fatal(err)
	}

	uri, err := svc.GetOAuthRedirect(provider)
	if err != nil {
		t.Fatal(err)
	}

	expected := "http://fake.oauth"
	if uri != expected {
		t.Fatalf("expected '%s', got '%s", expected, uri)
	}
}

func TestAuthorize(t *testing.T) {
	svc, err := auth.New()
	if err != nil {
		t.Fatal(err)
	}

	token, err := svc.Authorize(provider, "fake_code")
	if err != nil {
		t.Fatal(err)
	}

	expectedUsername := "fake"
	if token.Username != expectedUsername {
		t.Fatalf("expected '%s', got '%s", expectedUsername, token.Username)
	}

	if len(token.Token) == 0 {
		t.Fatalf("invalid token: %d", len(token.Token))
	}

	if time.Now().Unix() > token.Expiration {
		t.Fatalf("invalid expiration: %d", token.Expiration)
	}
}
