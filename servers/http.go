package servers

import (
	"context"
	"net/http"
)

type httpServer struct {
	srv *http.Server
}

// this way opens a http_server to listen actual connection, which I think not good
// but I will write a test for this
func (s httpServer) Run(h http.Handler) {
	s.srv.Handler = h
	go s.srv.ListenAndServe()
}

func (s httpServer) Stop() {
	s.srv.Shutdown(context.TODO())
}

func NewHttp(addr string) Server {
	return &httpServer{
		srv: &http.Server{
			Addr: addr,
		},
	}
}
