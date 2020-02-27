package httpServer

import (
	"context"
	"net/http"
)

type Server interface {
	Run(http.Handler)
	Stop()
}

type server struct {
	srv *http.Server
}

// this way opens a http_server to listen actual connection, which I think not good
// but I will write a test for this
func (s server) Run(h http.Handler) {
	s.srv.Handler = h
	go s.srv.ListenAndServe()
}

func (s server) Stop() {
	s.srv.Shutdown(context.TODO())
}

func New(addr string) Server {
	return &server{
		srv: &http.Server{
			Addr: addr,
		},
	}
}
