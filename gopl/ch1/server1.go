package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)                         // each request calls a handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // handler echoes the path component of the requester URL
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
