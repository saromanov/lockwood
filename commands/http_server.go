package commands

import (
   "net"
   "net/http"
   "log"
)

type HttpServer struct{
	Addr string
	
}

func (hs *HttpServer) Register() {
	http.HandleFunc("/v1/serfstat", hs.serfstat)
	http.HandleFunc("/v1/send", hs.send)
}

func (hs *HttpServer) Run() {
	server := http.Server {
		Addr: hs.Addr,
	}

	go server.ListenAndServe()
}