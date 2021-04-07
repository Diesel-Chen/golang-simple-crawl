package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
)

func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept error:%s", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
	return nil
}
