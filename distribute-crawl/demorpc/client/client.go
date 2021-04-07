package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClient(conn)
	var result float64
	err = client.Call("DemoRpcService.Div", struct{ A, B int }{25, 100}, &result)
	if err != nil {
		log.Println("call failed err:", err)
	} else {
		log.Println("result:", result)
	}

}
