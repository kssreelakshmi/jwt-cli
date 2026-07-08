package main

import (
	"strings"
	"testing"
)

func TestCheckExpiry_ValidToken(t *testing.T) {
	// exp: 2000000000 -> May 2033, far in the future
	payload := map[string]interface{}{
		"exp": float64(2000000000),
	}

	result := checkExpiry(payload)

	if !strings.Contains(result, "TOKEN VALID") {
		t.Errorf("expected TOKEN VALID, got: %s", result)
	}
}

func TestCheckExpiry_ExpiredToken(t *testing.T) {
	// exp: 1700000000 -> Nov 2023, in the past
	payload := map[string]interface{}{
		"exp": float64(1700000000),
	}

	result := checkExpiry(payload)

	if !strings.Contains(result, "TOKEN EXPIRED") {
		t.Errorf("expected TOKEN EXPIRED, got: %s", result)
	}
}

func TestDecodeSegment_ValidBase64(t *testing.T) {
	// base64url encoding of {"alg":"HS256","typ":"JWT"}
	segment := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

	result, err := decodeSegment(segment)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result["alg"] != "HS256" {
		t.Errorf("expected alg HS256, got: %v", result["alg"])
	}
}

func TestDecodeSegment_InvalidBase64(t *testing.T) {
	// "test" is not valid base64url-encoded JSON
	segment := "test"

	_, err := decodeSegment(segment)
	if err == nil {
		t.Error("expected an error for invalid base64/JSON, got none")
	}
}
