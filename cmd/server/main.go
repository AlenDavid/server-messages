package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New request at", time.Now(), "Request URI is: ", r.RequestURI)

		fmt.Fprintf(w, "Welcome to new server!")
	})

	fmt.Println("Server is ready on :5050")

	// listen to port
	http.ListenAndServe(":5050", nil)
}
