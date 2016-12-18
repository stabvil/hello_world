package main

import (
	"net/http"
	"github.com/gorilla/mux"	
	"encoding/json"
//	"fmt"
	
)



func index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "text/plain")
	
	slcB, _ := json.Marshal("ok")
	w.Write([]byte(slcB))
}




func main() {
	//var servStatus string
	//servStatus = "olala"
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/status", index)
	http.Handle("/", r)
	
	http.ListenAndServe(":8000", nil)
	
	//var input string
	//fmt.Scanln(&input)
}