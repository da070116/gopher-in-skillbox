package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skillbox-test/pkg"
	"syscall"
	"time"
)

const proxyAddr string = "localhost:9000"

var (
	counter            int    = 0
	firstInstanceHost  string = "http://localhost:8088/users/"
	secondInstanceHost string = "http://localhost:8087/users/"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osCall := <-c
		log.Printf("System call %+v\n", osCall)
		cancel()
	}()

	if err := runProxy(ctx); err != nil {
		log.Printf("failed to launch server :%+v\n", err)
	}

}

func runProxy(ctx context.Context) (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleQueries)

	proxySrv := &http.Server{Addr: proxyAddr, Handler: mux}
	go func() {
		if err = proxySrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	log.Println("Proxy is up at " + proxyAddr)

	<-ctx.Done()

	log.Printf("Server on %s stopped", proxyAddr)

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = proxySrv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("shutdown failed due to %s\n", err)
	}

	log.Printf("server %s exited correctly\n", proxyAddr)

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func handleQueries(w http.ResponseWriter, r *http.Request) {
	bodyContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer pkg.CloseReader(r.Body)

	if counter == 0 {
		if _, err := http.Post(firstInstanceHost, "text/json", bytes.NewBuffer(bodyContent)); err != nil {
			log.Fatalln(err)
		}

		counter++
		_, _ = w.Write([]byte("Data sent to " + firstInstanceHost))
		return
	}

	if _, err := http.Post(secondInstanceHost, "text/json", bytes.NewBuffer(bodyContent)); err != nil {
		log.Fatalln(err)
	}
	counter--
	_, _ = w.Write([]byte("Data sent to " + secondInstanceHost))
}
