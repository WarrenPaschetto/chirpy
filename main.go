package main

import (
	"net/http"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Create new http server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
