package producer

import (
	"github.com/Slowhigh/gogger/producer/infra/config"
	"github.com/Slowhigh/gogger/producer/internal/entity/interactor"
	"github.com/memphisdev/memphis.go"
)

type Producer struct {
	conn     *memphis.Conn
	producer *memphis.Producer
}

func NewProducer(conf *config.Config) (interactor.Producer, error) {
	conn, err := memphis.Connect(conf.Memphis.Host, conf.Memphis.UserName, memphis.Password(conf.Memphis.Password))
	if err != nil {
		return nil, err
	}

	producer, err := conn.CreateProducer(conf.Memphis.StationName, conf.Memphis.ProducerName)
	if err != nil {
		return nil, err
	}

	return &Producer{
		conn:     conn,
		producer: producer,
	}, nil
}

// Produce implements domain.MessageProducer.
func (mp *Producer) Produce(msg any) error {
	return mp.producer.Produce(msg)
}
