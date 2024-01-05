package config

import (
	"log/slog"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Http    Http    `mapstructure:"http"`
	Memphis Memphis `mapstructure:"memphis"`
}

type Http struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // "debug" or "release"
}

type Memphis struct {
	Host         string `mapstructure:"host"`
	UserName     string `mapstructure:"user_name"`
	Password     string `mapstructure:"password"`
	StationName  string `mapstructure:"station_name"`
	ProducerName string `mapstructure:"producer_name"`
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
