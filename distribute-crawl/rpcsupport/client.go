package rpcsupport

import (
	"net"
	"net/rpc"
)

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return rpc.NewClient(conn), nil
}
