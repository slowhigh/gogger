package interactor

import (
	"github.com/Slowhigh/gogger/producer/internal/entity/proto"
)

type Producer interface {
	ProduceAccessLog(msg *proto.AccessLog) error
	// ProduceJobLog(msg *proto.JobLog) error
}
