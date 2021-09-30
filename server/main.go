package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

var (
	port = "7000"
)

type HelloService struct{}

type SayHelloReply struct {
	Message string
}

func (s *HelloService) SayHello(text []byte, res *SayHelloReply) error {
	message := string(text)
	fmt.Printf("Mensagem recebida: %v\n", message)
	*res = SayHelloReply{"Olá " + message}
	return nil
}

func main() {
	address, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	server, err := net.ListenTCP("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(HelloService)
	rpc.Register(listener)
	fmt.Println("Listen Server RPC on port: " + port)
	rpc.Accept(server)
}
