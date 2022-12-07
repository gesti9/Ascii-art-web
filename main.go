package main

import (
	"fmt"
	"net/http"
	home "workspace/controler"
)

func handleRequest() {
	fmt.Println("GO to: http://127.0.0.1:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home.Home)
	mux.HandleFunc("/ascii-art", home.AsciiArt)
	http.ListenAndServe(":8080", mux)
}

func main() {
	handleRequest()
}
