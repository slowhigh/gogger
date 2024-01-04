package main

import (
	"log"

	"github.com/memphisdev/memphis.go"
)

type Server interface {
	Run() error
}

func main() {
	conn, err := memphis.Connect("localhost", "producer_1", memphis.Password("#B8T2oA-mZ"), memphis.AccountId(1))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	var server Server

	server, err = InitializeServer(conn)
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Run(); err != nil {
		log.Fatalln()
	}
}
