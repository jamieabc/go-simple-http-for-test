package server

import (
	"io"
	"net/http"
)

type Server interface {
	Run()
	Stop()
}

type server struct {
	addr string
}

// this way opens a server to listen actual connection, which I think not good
// but I will write a test for this
func (s server) Run() {
	handler := func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "hello, world!\n")
	}
	http.HandleFunc("/hello", handler)
	go http.ListenAndServe(s.addr, nil)
}

func (s server) Stop() {
	panic("implement me")
}

func New(addr string) Server {
	return &server{
		addr: addr,
	}
}
