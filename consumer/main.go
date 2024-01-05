package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Slowhigh/gogger/consumer/internal/entity"
	"github.com/Slowhigh/gogger/consumer/proto"
	"github.com/memphisdev/memphis.go"
	protobuf "google.golang.org/protobuf/proto"
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
			var accessLogPb proto.AccessLog
			err = protobuf.Unmarshal(msg.Data(), &accessLogPb)
			if err != nil {
				panic(err)
			}

			accessLog := entity.AccessLog{
				Timestamp:    accessLogPb.Timestamp.AsTime(),
				IsNormalMode: accessLogPb.IsNormalMode,
				IsLogin:      accessLogPb.IsLogin,
				UserName:     accessLogPb.UserName,
				DeviceName:   accessLogPb.DeviceName,
				Ip:           accessLogPb.Ip,
			}

			fmt.Printf("[Consume] %+v\n", accessLog)
			msg.Ack()
		}
	}

	ctx := context.Background()
	consumer.SetContext(ctx)
	consumer.Consume(handler)
	time.Sleep(30 * time.Minute)
}
