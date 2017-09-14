package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Web Server from ODX Registry!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}
