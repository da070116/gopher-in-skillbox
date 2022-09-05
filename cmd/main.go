package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skillbox-test/pkg"
	"syscall"
	"time"
)

// main - entry point
func main() {

	if _, err := os.Stat("")

	// get args from console to define port and host
	port := flag.Int("port", 8080, "port to launch server")
	host := flag.String("host", "localhost", "host to launch server")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-c
		log.Printf("System call %+v\n", osCall)
		cancel()
	}()

	if err := runServer(ctx, *host, *port); err != nil {
		log.Printf("failed to launch server :%+v\n", err)
	}
}

// runServer - configure and run server with graceful shutdown
func runServer(ctx context.Context, host string, port int) (err error) {

	mux := http.NewServeMux()
	service := pkg.Service{Storage: make(map[int]pkg.User)}
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		pkg.HandlersManager(w, r, &service)
	})

	addr := fmt.Sprintf("%s:%d", host, port)

	srv := &http.Server{Addr: addr, Handler: mux}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	log.Println("Launch server on " + addr)

	<-ctx.Done()

	log.Printf("Server on %s stopped", addr)

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("shutdown failed due to %s\n", err)
	}

	log.Printf("server %s exited correctly\n", addr)

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
