package listener

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener struct {
	rcvr any
}

func New(rcvr any) *Listener {
	return &Listener{rcvr}
}

func (l *Listener) Listen(port string) error {
	err := rpc.Register(l.rcvr)
	if err != nil {
		return err
	}

	log.Println("Starting RPC server on port ", port)
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
