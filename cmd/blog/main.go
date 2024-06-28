package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luckyman9339/blog/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/signup", handlers.SignupHandler).Methods("GET", "POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	router.HandleFunc("/posts", handlers.A)
}
