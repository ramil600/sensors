package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Dbname string `yaml:"db_name"`
	PublishDSN string `yaml:"publish_dsn"`
}

func NewConfig() *Config {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		PublishDSN: viper.GetString("publish_dsn"),
		Dbname: viper.GetString("db_name"),
	}
}