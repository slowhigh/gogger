package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/memphisdev/memphis.go"
)

func main() {
	conn, err := memphis.Connect("localhost", "consumer1", memphis.Password("587#h%@lW#"), memphis.AccountId(1))
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
			fmt.Printf("Fetch failed: %v", err)
			return
		}

		for _, msg := range msgs {
			fmt.Println(string(msg.Data()))
			msg.Ack()
		}
	}

	ctx := context.Background()
	// ctx = context.WithValue(ctx, "key", "value")
	consumer.SetContext(ctx)
	consumer.Consume(handler)
	// The program will close the connection after 30 seconds,
	// the message handler may be called after the connection closed
	// so the handler may receive a timeout error
	time.Sleep(30 * time.Minute)
}
