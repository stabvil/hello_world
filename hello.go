package main

import (
    
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Only matches if domain is "www.example.com".
	r.Host("www.example.com")
	// Matches a dynamic subdomain.
	
    /*r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/products", ProductsHandler)
    r.HandleFunc("/articles", ArticlesHandler)
    http.Handle("/", r)*/
}