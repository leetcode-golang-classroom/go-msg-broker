package common

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RABBITMQ_URL string `mapstructure:"RABBITMQ_URL"`
}

var C *Config

func init() {
	log.Printf("init config")
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if err != nil {
		failOnError(err, "Failed to read config")
	}
	v.AutomaticEnv()

	err = v.Unmarshal(&C)
	if err != nil {
		failOnError(err, "Failed to read environment")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
