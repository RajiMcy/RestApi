package main

import (
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found \n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1 \n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
