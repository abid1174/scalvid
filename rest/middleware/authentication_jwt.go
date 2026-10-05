package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticationJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// extract jwt token from request header
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Unauthorised", http.StatusUnauthorized)
			return
		}

		authHeaderArr := strings.Split(authHeader, " ")

		if len(authHeaderArr) != 2 {
			http.Error(w, "Unauthorised", http.StatusUnauthorized)
			return
		}

		accessToken := authHeaderArr[1]

		// parse accesstoken
		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorised", http.StatusUnauthorized)
			return
		}

		tokenHeader := tokenParts[0]
		tokenPayload := tokenParts[1]
		tokenSignature := tokenParts[2]

		// Generate token new signature
		tokenHeaderAndPayload := tokenHeader + "." + tokenPayload
		jwtSecret := m.config.JwtSecretKey

		byteArrTokenHeaderAndPayload := []byte(tokenHeaderAndPayload)
		byteArrJwtSecret := []byte(jwtSecret)

		h := hmac.New(sha256.New, byteArrJwtSecret)
		h.Write(byteArrTokenHeaderAndPayload)
		hash := h.Sum(nil)
		newSignature := base64Encode(hash)

		// check token signature with newly created signature
		if tokenSignature != newSignature {
			http.Error(w, "Unauthorised", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
