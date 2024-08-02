package config

import (
	"github.com/spf13/viper"
)

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`

	ServerPort int `mapstructure:"SERVER_PORT"`

	DBUrl string `mapstructure:"DB_URL"`
}

func (c Config) IsProduction() bool {
	return c.Environment == EnvProduction
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
