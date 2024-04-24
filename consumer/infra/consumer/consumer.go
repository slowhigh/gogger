package consumer

import (
	"context"
	"log/slog"
	"time"

	"github.com/Slowhigh/gogger/consumer/infra/config"
	"github.com/Slowhigh/gogger/consumer/infra/consumer/handler"
	"github.com/Slowhigh/gogger/consumer/internal/adapter/controller"
	"github.com/memphisdev/memphis.go"
)

type Consumer struct {
	conn              *memphis.Conn
	accessLogConsumer *memphis.Consumer
	ctrl              controller.Controller
}

func NewConsumer(conf *config.Config, sc controller.Controller) (*Consumer, error) {
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
		ctrl:              sc,
	}, nil
}

func (c Consumer) Run() error {
	err := c.accessLogConsumer.Consume(
		func(m []*memphis.Msg, err error, ctx context.Context) {
			handler.AccessLogHandler(c.ctrl.ConsumeAccessLog, m, err, ctx)
		},
	)
	if err != nil {
		return err
	}

	slog.Info("start consuming messages")
	return nil
}

func (c Consumer) Stop() {
	c.accessLogConsumer.StopConsume()
	slog.Info("stop consuming messages")
}
