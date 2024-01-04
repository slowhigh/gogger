package producer

import (
	"github.com/Slowhigh/gogger/producer/internal/entity/interactor"
	"github.com/memphisdev/memphis.go"
)

type Producer struct {
	producer *memphis.Producer
}

func NewProducer(conn *memphis.Conn) (interactor.Producer, error) {
	// TODO: `stationName`, `producerName` environment variable.
	p, err := conn.CreateProducer("default", "pro-1")
	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: p,
	}, nil
}

// Produce implements domain.MessageProducer.
func (mp *Producer) Produce(msg any) error {
	return mp.producer.Produce(msg)
}
