package main

import (
	"io"
	"log"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	html := "<html><form metho=\"POST\" action=\"/upload\"  enctype=\"multipart/form-data\"> Choose an image to upload: <input name=\"image\" type=\"file\" /><input type=\"submit\" value=\"Upload\" /></form></html>"
	if r.Method == "GET" {
		io.WriteString(w, html)
		return
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
