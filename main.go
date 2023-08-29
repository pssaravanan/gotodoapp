package main

import (
	"io"
	"net/http"
)

func elbCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func main(){
	http.HandleFunc("/ping", elbCheck)
	http.ListenAndServe(":3000", nil)
}