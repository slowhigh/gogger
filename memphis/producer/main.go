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

	for i := 0; i < 1000000; i++ {
		message := Message{
			Message: fmt.Sprintf("msg - %d", i),
		}

		binary, err := json.Marshal(message)
		if err != nil {
			panic(err)
		}

		err = p.Produce(binary, memphis.AsyncProduce())
		if err != nil {
			panic(err)
		}

		fmt.Printf("[Publish] %s\n", message.Message)
	}
}
