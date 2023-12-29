package infrastructure

import (
	"github.com/Slowhigh/gogger/producer/internal/entity/interactor"
	"github.com/memphisdev/memphis.go"
)

type MessageProducer struct {
	producer *memphis.Producer
}

func NewMessageProducer(conn *memphis.Conn) (interactor.MessageProducer, error) {
	// TODO: `stationName`, `producerName` environment variable.
	p, err := conn.CreateProducer("default", "pro-1")
	if err != nil {
		return nil, err
	}

	return &MessageProducer{
		producer: p,
	}, nil
}

// Produce implements domain.MessageProducer.
func (mp *MessageProducer) Produce(msg any) error {
	return mp.producer.Produce(msg)
}
