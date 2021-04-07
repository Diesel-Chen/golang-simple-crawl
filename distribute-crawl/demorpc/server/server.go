package main

import (
	"golang-simple-crawl/distribute-crawl/demorpc"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.Register(demorpc.DemoRpcService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept error:%s", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
