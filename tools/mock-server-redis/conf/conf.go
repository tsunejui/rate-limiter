package conf

import (
	"os"
)

type Config struct {
	ListenPort    string
	RedisAddress  string
	RedisPassword string
}

const (
	envListenPort    string = "LISTEN_PORT"
	envRedisAddress  string = "REDIS_ADDRESS"
	envRedisPassword string = "REDIS_PASSWORD"
)

var config = &Config{}

func init() {
	config.ListenPort = os.Getenv(envListenPort)
	config.RedisAddress = os.Getenv(envRedisAddress)
	config.RedisPassword = os.Getenv(envRedisPassword)
}

func GetConfig() *Config {
	return config
}
