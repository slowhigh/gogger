package interactor

import "github.com/Slowhigh/gogger/consumer/internal/entity"

type AccessLog interface {
	Create(*entity.AccessLog) error
}
