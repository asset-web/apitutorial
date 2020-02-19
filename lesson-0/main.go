package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	log.Println("service started on :8080")
	http.ListenAndServe(":8080", nil)
	log.Println("Am I ever run??? Nope!")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Log all the things")
	w.Write([]byte("Hello, world!"))
}
