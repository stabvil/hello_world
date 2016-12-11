package main

import (
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte("Hello World!!!"))
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":80", nil)
	var input string
	fmt.Scanln(&input)
}