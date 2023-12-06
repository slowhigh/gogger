package main

import (
	"fmt"
	"os"

	"github.com/Slowhigh/gogger/producer/proto"
	"github.com/memphisdev/memphis.go"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	conn, err := memphis.Connect("localhost", "producer_1", memphis.Password("#B8T2oA-mZ"), memphis.AccountId(1))
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
		info := proto.Info{
			Msg: fmt.Sprintf("msg - %d", i),
		}

		err = p.Produce(&info, memphis.AsyncProduce())
		if err != nil {
			panic(err)
		}

		fmt.Printf("[Produce] %s\n", info.Msg)
	}
}
