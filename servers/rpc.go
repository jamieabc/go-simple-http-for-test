package servers

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Calculator struct{}

func (c Calculator) Add(arg RpcArg, reply *int) error {
	fmt.Println("add")
	*reply = arg.A + arg.B
	return nil
}

type rpcServer struct {
	port     string
	listener net.Listener
	srv      *rpc.Server
}

func (r *rpcServer) Run(_ http.Handler) {
	r.srv = rpc.NewServer()

	_ = r.srv.Register(new(Calculator))
	r.srv.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	l, _ := net.Listen("tcp", r.port)
	go http.Serve(l, nil)
}

func (r rpcServer) Stop() {
}

type RpcArg struct {
	A, B int
}

func NewRpc(port string) Server {
	return &rpcServer{
		port: port,
	}
}
