package server

import (
	"net/http"
)

type Server interface {
	Run(http.Handler)
	Stop()
}

type server struct {
	addr string
}

// this way opens a server to listen actual connection, which I think not good
// but I will write a test for this
func (s server) Run(h http.Handler) {
	go http.ListenAndServe(s.addr, h)
}

func (s server) Stop() {
	panic("implement me")
}

func New(addr string) Server {
	return &server{
		addr: addr,
	}
}
