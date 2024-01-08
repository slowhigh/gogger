package producer

import (
	"github.com/Slowhigh/gogger/producer/infra/config"
	"github.com/Slowhigh/gogger/producer/internal/entity/proto"
	"github.com/Slowhigh/gogger/producer/internal/usecase/interactor"
	"github.com/memphisdev/memphis.go"
)

type Producer struct {
	conn              *memphis.Conn
	accessLogProducer *memphis.Producer
	// jobLogProducer *memphis.Producer
}

func NewProducer(conf *config.Config) (interactor.Producer, error) {
	conn, err := memphis.Connect(conf.Memphis.Host, conf.Memphis.UserName, memphis.Password(conf.Memphis.Password))
	if err != nil {
		return nil, err
	}

	accessLogProducer, err := conn.CreateProducer(conf.Memphis.AccessMessageStationName, conf.Memphis.ProducerName)
	if err != nil {
		return nil, err
	}

	// jobLogProducer, err := conn.CreateProducer(conf.Memphis.JobMessageStationName, conf.Memphis.ProducerName)
	// if err != nil {
	// 	return nil, err
	// }

	return &Producer{
		conn:              conn,
		accessLogProducer: accessLogProducer,
		//jobLogProducer: jobLogProducer,
	}, nil
}

// ProduceAccessLog implements interactor.Producer.
func (mp *Producer) ProduceAccessLog(msg *proto.AccessLog) error {
	return mp.accessLogProducer.Produce(msg)
}

// func (mp *Producer) ProduceJobLog(msg *proto.JobLog) error {
// 	return mp.jobLogProducer.Produce(msg)
// }
