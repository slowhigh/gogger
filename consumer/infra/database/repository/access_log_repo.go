package repository

import (
	"github.com/Slowhigh/gogger/consumer/infra/database"
	"github.com/Slowhigh/gogger/consumer/internal/entity"
	"github.com/Slowhigh/gogger/consumer/internal/usecase/interactor"
)

type AccessLog struct {
	db *database.Database
}

func NewAccessLog(db *database.Database) (interactor.AccessLog, error) {
	if err := db.AutoMigrate(&entity.AccessLog{}); err != nil {
		return nil, err
	}

	return &AccessLog{
		db: db,
	}, nil
}

// Create implements interactor.AccessLog.
func (al *AccessLog) Create(accessLog *entity.AccessLog) error {
	err := al.db.Create(&accessLog).Error
	if err != nil {
		return err
	}

	return nil
}
