package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

// ============================================================================
// JWT Data Structures
// ============================================================================

// Header represents the JWT header containing algorithm and token type.
// This is the first part of the JWT (before the first dot).
type Header struct {
	Alg string `json:"alg"` // Algorithm used for signing (e.g., "HS256")
	Typ string `json:"typ"` // Token type (always "JWT")
}

// Payload represents the JWT payload containing user claims and metadata.
// This is the second part of the JWT (between the two dots).
type Payload struct {
	Sub   string `json:"sub"`   // Subject (typically user ID)
	Email string `json:"email"` // User's email address
	Role  string `json:"role"`  // User's role (e.g., "admin", "user")
	Exp   int64  `json:"exp"`   // Expiration time (Unix timestamp)
	Iat   int64  `json:"iat"`   // Issued at time (Unix timestamp)
}

// ============================================================================
// JWT Generation
// ============================================================================

// GenerateJWT creates a signed JWT token using HMAC-SHA256 algorithm.
// The token consists of three parts separated by dots: header.payload.signature
//
// Parameters:
//   - payload: The JWT payload containing user information and claims
//   - secretKey: The secret key used to sign the token
//
// Returns:
//   - string: The complete JWT token
//   - error: Any error that occurred during token generation
func GenerateJWT(payload Payload, secretKey string) (string, error) {
	// Step 1: Create the JWT header
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// Step 2: Set token timestamps
	payload.Iat = time.Now().Unix()
	payload.Exp = time.Now().Add(1 * time.Hour).Unix()

	// Step 3: Serialize header and payload to JSON
	byteArrayHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	byteArrayPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Step 4: Encode header and payload to base64
	base64Header := base64Encode(byteArrayHeader)
	base64Payload := base64Encode(byteArrayPayload)

	// Step 5: Combine header and payload with a dot separator
	base64HeaderAndPayload := base64Header + "." + base64Payload

	// Step 6: Generate HMAC-SHA256 signature
	byteArraySecretKey := []byte(secretKey)
	byteArrayHeaderAndPayload := []byte(base64HeaderAndPayload)

	h := hmac.New(sha256.New, byteArraySecretKey)
	h.Write(byteArrayHeaderAndPayload)
	signature := h.Sum(nil)
	base64Signature := base64Encode(signature)

	// Step 7: Combine all three parts to form the complete JWT
	jwtToken := base64HeaderAndPayload + "." + base64Signature
	return jwtToken, nil
}

// ============================================================================
// Helper Functions
// ============================================================================

// base64Encode encodes data to URL-safe base64 without padding.
// This is the standard encoding format required by JWT specification.
func base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
