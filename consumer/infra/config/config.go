package config

import (
	"log/slog"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Memphis  Memphis  `mapstructure:"memphis"`
	Database Database `mapstructure:"database"`
}

type Memphis struct {
	Host                     string `mapstructure:"host"`
	UserName                 string `mapstructure:"user_name"`
	Password                 string `mapstructure:"password"`
	AccessMessageStationName string `mapstructure:"access_message_station_name"`
	ConsumerName             string `mapstructure:"consumer_name"`
	ConsumerGroupName        string `mapstructure:"consumer_group_name"`
	PullInterval             int    `mapstructure:"pull_interval"`
	BatchSize                int    `mapstructure:"batch_size"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}

func NewConfig() (*Config, error) {
	_, b, _, _ := runtime.Caller(0)
	configDirPath := path.Join(path.Dir(b))

	conf := Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDirPath)

	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("Read config file.", "err", err)
		return nil, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		slog.Error("Unmarshal config file.", "err", err)
		return nil, err
	}

	return &conf, nil
}
