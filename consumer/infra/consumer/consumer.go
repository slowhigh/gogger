package consumer

import (
	"context"
	"log/slog"
	"time"

	"github.com/Slowhigh/gogger/consumer/infra/config"
	"github.com/Slowhigh/gogger/consumer/internal/adapter/controller/sub"
	"github.com/Slowhigh/gogger/consumer/internal/entity/proto"
	"github.com/memphisdev/memphis.go"
	protobuf "google.golang.org/protobuf/proto"
)

type Consumer struct {
	conn              *memphis.Conn
	accessLogConsumer *memphis.Consumer
	subCtrl           sub.Controller
}

func NewConsumer(conf *config.Config, sc sub.Controller) (*Consumer, error) {
	conn, err := memphis.Connect(conf.Memphis.Host, conf.Memphis.UserName, memphis.Password(conf.Memphis.Password))
	if err != nil {
		return nil, err
	}

	alc, err := conn.CreateConsumer(
		conf.Memphis.AccessMessageStationName,
		conf.Memphis.ConsumerName,
		memphis.ConsumerGroup(conf.Memphis.ConsumerGroupName),
		memphis.PullInterval(time.Duration(conf.Memphis.PullInterval)*time.Millisecond),
		memphis.BatchSize(conf.Memphis.BatchSize),
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn:              conn,
		accessLogConsumer: alc,
		subCtrl:           sc,
	}, nil
}

func (c Consumer) Run() error {
	accessLogHandler := func(msgs []*memphis.Msg, err error, ctx context.Context) {
		if err != nil {
			return
		}

		for _, msg := range msgs {
			var accessLogPb proto.AccessLog
			err = protobuf.Unmarshal(msg.Data(), &accessLogPb)
			if err != nil {
				panic(err)
			}

			if c.subCtrl.ConsumeAccessLog(&accessLogPb) {
				msg.Ack()
			}
		}
	}

	if err := c.accessLogConsumer.Consume(accessLogHandler); err != nil {
		return err
	}

	slog.Info("start consuming messages")
	return nil
}

func (c Consumer) Stop() {
	c.accessLogConsumer.StopConsume()
	slog.Info("stop consuming messages")
}
