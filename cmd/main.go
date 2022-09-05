package main

import (
	"log"
	"net/http"
	"skillbox-test/pkg"
)

// main - entry point
func main() {

	mux := http.NewServeMux()
	service := pkg.Service{Storage: make(map[int]pkg.User)}

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		pkg.HandlersManager(w, r, &service)
	})

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
