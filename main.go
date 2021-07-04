package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World main")
}
func main() {
	server := &http.Server{
		Addr: "0.0.0.0:8888",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
