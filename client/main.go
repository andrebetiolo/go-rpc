package main

import (
	"bufio"
	"log"
	"net/rpc"
	"os"
)

var (
	port = "7000"
)

type SayHelloReply struct {
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	stdin := bufio.NewReader(os.Stdin)
	for {
		message, _, err := stdin.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		var reply SayHelloReply
		err = client.Call("HelloService.SayHello", message, &reply)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Mensagem recebida do servidor: %v", reply.Message)
	}
}
