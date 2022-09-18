package gopher_in_skillbox

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	inst *http.Server
}

// Run - start server
func (s *Server) Run(handler http.Handler, host string, port int) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	s.inst = &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        handler,
	}
	return s.inst.ListenAndServe()
}

// Shutdown - stop server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.inst.Shutdown(ctx)
}
