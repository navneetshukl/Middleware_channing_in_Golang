package main

import (
	"fmt"
	"go_modules/middleware"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index Handler")
	fmt.Fprintf(w, "Welcome")
}
func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing message Handler")
	fmt.Fprintf(w, "HTTP Middleware is awesome")
}
func iconHandler(w http.ResponseWriter, r *http.Request) {
}
func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	http.Handle("/", middleware.MiddlewareFirst(middleware.MiddlewareSecond(http.HandlerFunc(index))))
	http.Handle("/message", middleware.MiddlewareFirst(middleware.MiddlewareSecond(http.HandlerFunc(message))))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
