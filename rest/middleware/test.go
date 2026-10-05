package middleware

import (
	"log"
	"net/http"
)

func Test(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Test")
		next.ServeHTTP(w, r)
	})
}
