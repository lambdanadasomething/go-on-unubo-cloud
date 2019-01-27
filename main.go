package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to Go on Unubo Cloud")
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
