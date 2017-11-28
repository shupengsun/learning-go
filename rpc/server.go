package main

import (
	"learning-go/rpc/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	arith := new(service.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	go http.Serve(l, nil)

	time.Sleep(100 * time.Second)
}
