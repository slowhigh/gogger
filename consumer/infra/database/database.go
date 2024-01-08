package database

import (
	"fmt"

	"github.com/Slowhigh/gogger/consumer/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(conf *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.UserName,
		conf.Database.Password,
		conf.Database.DbName,
	)

	postgres, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: postgres,
	}, nil
}
