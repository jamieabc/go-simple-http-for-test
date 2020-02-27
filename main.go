package main

import (
	"github.com/jamieabc/go-simple-http-for-test/servers"
	"os"
	"os/signal"
)

func main() {
	//srvHttp := servers.NewHttp(":5566")
	//srvHttp.Run(handlers.New())

	srvRpc := servers.NewRpc(":9453")
	srvRpc.Run(nil)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
