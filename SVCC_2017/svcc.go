package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9090", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "SVCC was fun!")
}
