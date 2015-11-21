package lockwood

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"runtime"
)

//Initialize RPC server

var (
	ErrRegister = errors.New("Error of function registration")
)

type RPCServer struct {
	listen *net.Listener
}

func InitRPCServer() *RPCServer {
	runtime.GOMAXPROCS(runtime.NumCPU())
	return &RPCServer{}
}

//Close provides stopping of RPC server
func (rpcs *RPCServer) Close() {
	rpcs.Close()
}

// register provides registration of enpoint
func (rpcs *RPCServer) register(item interface{}) error {
	if item == nil {
		return ErrRegister
	}
	rpc.Register(item)
}

// start of RPC server
func (rpcs *RPCServer) start(addr string) {
	listen, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpcs.listen = listen

	rpc.Accept(rpcs.listen)
}
