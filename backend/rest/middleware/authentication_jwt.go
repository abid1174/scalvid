package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"scalvid/utils"
	"strings"
	"time"
)

func (m *Middlewares) AuthenticationJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// extract jwt token from "Authorization: Bearer <token>" header
		authHeader := r.Header.Get("Authorization")

		scheme, accessToken, found := strings.Cut(authHeader, " ")
		if !found || !strings.EqualFold(scheme, "Bearer") || accessToken == "" {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorised")
			return
		}

		// parse accesstoken
		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorised")
			return
		}

		tokenHeader := tokenParts[0]
		tokenPayload := tokenParts[1]
		tokenSignature := tokenParts[2]

		// Generate token new signature
		tokenHeaderAndPayload := tokenHeader + "." + tokenPayload
		jwtSecret := m.config.JwtSecretKey

		h := hmac.New(sha256.New, []byte(jwtSecret))
		h.Write([]byte(tokenHeaderAndPayload))
		newSignature := base64Encode(h.Sum(nil))

		// check token signature with newly created signature in constant time
		if !hmac.Equal([]byte(tokenSignature), []byte(newSignature)) {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorised")
			return
		}

		// reject expired tokens
		payloadJSON, err := base64.RawURLEncoding.DecodeString(tokenPayload)
		if err != nil {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorised")
			return
		}

		var payload utils.Payload
		if err := json.Unmarshal(payloadJSON, &payload); err != nil {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorised")
			return
		}

		if payload.Exp == 0 || time.Now().Unix() >= payload.Exp {
			utils.SendError(w, http.StatusUnauthorized, "Token expired")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
