package middleware

import (
	"log"
	"net/http"
)

func MiddlewareFirst(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareFirst - Before Handler")
		next.ServeHTTP(w, r)
		log.Println("MiddlewareFirst - After Handler")
	})
}

func MiddlewareSecond(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddlewareSecond - Before Handler")
		if r.URL.Path == "/message" {
			if r.URL.Query().Get("password") == "pass123" {
				log.Println("Authorized to the system")
				next.ServeHTTP(w, r)
			} else {
				log.Println("Failed to authorize to the system")
				return
			}
		} else {
			next.ServeHTTP(w, r)
		}
		log.Println("MiddlewareSecond - After Handler")
	})
}
