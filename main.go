package main

import (
	"bp-auth-interceptor/authentication"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/verify", authentication.Verify)
	mux.HandleFunc("/create-token", authentication.CreateToken)

	serverStartLog()
	if err := http.ListenAndServe(":8200", mux); err != nil {
		log.Fatal(err)
	}
}

func serverStartLog() {
	log.Println("Starting server at port 8200")
}
