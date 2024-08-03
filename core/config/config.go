package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var conf Config

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`

	ServerPort int `mapstructure:"SERVER_PORT"`

	DBSource string `mapstructure:"DB_SOURCE"`

	RedisSource string `mapstructure:"REDIS_SOURCE"`
}

func (c Config) IsProduction() bool {
	return c.Environment == EnvProduction
}

func GetConfig() Config {
	return conf
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

func Init(path string) {
	var err error
	conf, err = LoadConfig(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env conf: %s\n", err)
		os.Exit(1)
	}
}
