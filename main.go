package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jwt-cli <token>")
		os.Exit(1)
	}

	token := os.Args[1]

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Error: invalid JWT format (expected 3 parts separated by '.')")
		os.Exit(1)
	}

	headerJSON, err := decodeSegment(parts[0])
	if err != nil {
		fmt.Println("Error decoding header:", err)
		os.Exit(1)
	}

	payloadJSON, err := decodeSegment(parts[1])
	if err != nil {
		fmt.Println("Error decoding payload:", err)
		os.Exit(1)
	}

	fmt.Println("=== HEADER ===")
	printPretty(headerJSON)

	fmt.Println("=== PAYLOAD ===")
	printPretty(payloadJSON)

	fmt.Println("=== STATUS ===")
	fmt.Println(checkExpiry(payloadJSON))
}

// decodeSegment base64url-decodes a single JWT segment and returns it as a map
func decodeSegment(segment string) (map[string]interface{}, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(decoded, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// printPretty pretty-prints a map as indented JSON
func printPretty(data map[string]interface{}) {
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(pretty))
}

// checkExpiry looks for the 'exp' claim in the payload and reports whether the token is expired
func checkExpiry(payload map[string]interface{}) string {
	expRaw, ok := payload["exp"]
	if !ok {
		return "NO EXPIRY CLAIM FOUND (token has no 'exp' field)"
	}

	expFloat, ok := expRaw.(float64)
	if !ok {
		return "INVALID EXPIRY CLAIM (exp is not a number)"
	}

	expTime := time.Unix(int64(expFloat), 0)

	if time.Now().After(expTime) {
		return fmt.Sprintf("TOKEN EXPIRED (expired at: %s)", expTime.Local())
	}
	return fmt.Sprintf("TOKEN VALID (expires at: %s)", expTime.Local())
}
