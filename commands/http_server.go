package commands

import (
   "net"
   "net/http"
   "log"
)

type HttpServer struct{
	Addr string

}

// Register provides registration of all handlers
func (hs *HttpServer) Register() {
	http.HandleFunc("/v1/serfstat", hs.serfstat)
	http.HandleFunc("/v1/send", hs.send)
}

// Run provides starting of server
func (hs *HttpServer) Run() {
	server := http.Server {
		Addr: hs.Addr,
	}

	go server.ListenAndServe()
}