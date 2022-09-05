package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"skillbox-test/pkg"
)

// main - entry point
func main() {
	port := flag.Int("port", 8080, "port to launch server")
	host := flag.String("host", "localhost", "host to launch server")

	flag.Parse()

	mux := http.NewServeMux()
	service := pkg.Service{Storage: make(map[int]pkg.User)}

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		pkg.HandlersManager(w, r, &service)
	})
	addr := fmt.Sprintf("%s:%d", *host, *port)

	println("Launch server on " + addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalln(err)
	}
}
