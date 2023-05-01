package config

import (
	"github.com/spf13/viper"
)

type (
	postgres struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Port     string `json:"port"`
	}

	AppConfig struct {
		Port         string   `json:"port"`
		TaskInterval int      `json:"task_interval" mapstructure:"task_interval"`
		Postgres     postgres `json:"postgres"`
	}
)

func New() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("internal/config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	ac := AppConfig{}
	if err := viper.Unmarshal(&ac); err != nil {
		return nil, err
	}

	return &ac, nil
}
