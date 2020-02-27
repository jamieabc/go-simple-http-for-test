package main

import (
	"github.com/jamieabc/go-simple-http-for-test/handlers"
	"github.com/jamieabc/go-simple-http-for-test/server"
	"os"
	"os/signal"
)

func main() {
	srv := server.New(":5566")
	srv.Run(handlers.New())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
