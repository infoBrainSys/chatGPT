package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Model       string
	MaxTokens   string
	Temperature float64
	TopP        int
	N           int
	Key         string
}

func (c *Config) GetCfg() (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigFile("config/config.json")
	if err := v.ReadInConfig(); err != nil {
		log.Println("Read In Config err :", err)
		return nil, err
	}
	return
}
