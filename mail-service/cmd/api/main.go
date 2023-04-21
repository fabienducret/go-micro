package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"strconv"
)

type Config struct {
	Mailer Mail
}

const rpcPort = "5001"

func main() {
	app := Config{
		Mailer: createMail(),
	}

	app.startServer()
}

func (app *Config) startServer() {
	rpcServer := new(RPCServer)
	rpcServer.Mailer = app.Mailer

	_ = rpc.Register(rpcServer)

	log.Println("Starting RPC server on port ", rpcPort)
	listen, _ := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
}

func createMail() Mail {
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("FROM_NAME"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
	}

	return m
}
