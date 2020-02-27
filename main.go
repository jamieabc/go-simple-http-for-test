package main

import (
	"github.com/jamieabc/go-simple-http-for-test/handlers"
	httpServer "github.com/jamieabc/go-simple-http-for-test/http_server"
	"os"
	"os/signal"
)

func main() {
	srv := httpServer.New(":5566")
	srv.Run(handlers.New())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
