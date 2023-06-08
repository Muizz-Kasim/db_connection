package main

import (
	"fmt"
	"net/http"
	"gitlab.com/go_progresif/config"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {

	// HTTP router
	// Handle url path requests
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	// Establish Database connection
	config.Connect()
	
	// API Endpoints??? - RESTful API endpoint
	// To be developed

	// Start local http server @ PORT:3000
	fmt.Println("Server Started at PORT:3000")
	http.ListenAndServe(":3000", nil)
}
