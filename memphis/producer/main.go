package main

import (
	"encoding/json"
	"fmt"
	"github.com/memphisdev/memphis.go"
	"os"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	conn, err := memphis.Connect("localhost", "producer1", memphis.Password("#B8T2oA-mZ"), memphis.AccountId(1))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	p, err := conn.CreateProducer("default", "pro-1")
	if err != nil {
		fmt.Printf("Producer failed: %v", err)
		os.Exit(1)
	}

	message := Message{
		Message: "new message",
	}

	binary, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	err = p.Produce(binary, memphis.AsyncProduce())
	if err != nil {
		panic(err)
	}
}
