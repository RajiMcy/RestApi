package main

import (
	"GoApi/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/users/", handlers.UsersRouter)
	mux.HandleFunc("/users", handlers.UsersRouter)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
