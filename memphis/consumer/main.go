package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/memphisdev/memphis.go"
)

func main() {
	conn, err := memphis.Connect("localhost", "consumer_1", memphis.Password("587#h%@lW#"), memphis.AccountId(1))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	consumer, err := conn.CreateConsumer("default", "con-1", memphis.ConsumerGroup("con"), memphis.PullInterval(1*time.Nanosecond), memphis.BatchSize(100))
	if err != nil {
		fmt.Printf("Consumer creation failed: %v", err)
		os.Exit(1)
	}

	handler := func(msgs []*memphis.Msg, err error, ctx context.Context) {
		if err != nil {
			panic(err)
		}

		for _, msg := range msgs {
			data, err := msg.DataDeserialized()

			if err != nil {
				panic(err)
			}
			fmt.Printf("[Consume] %+v\n", data)
			msg.Ack()
		}
	}

	ctx := context.Background()
	consumer.SetContext(ctx)
	consumer.Consume(handler)
	time.Sleep(30 * time.Minute)
}
