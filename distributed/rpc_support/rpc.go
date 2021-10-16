package rpc_support

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRPC(host string, service interface{}) error {
	rpc.Register(service)
	log.Println("Save server running...")
	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v\n", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	return jsonrpc.NewClient(conn), nil
}