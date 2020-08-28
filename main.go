package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Server runing on port 8080")
	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CI/CD"))
}
