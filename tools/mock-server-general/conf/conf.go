package conf

import (
	"os"
)

type Config struct {
	ListenPort string
}

const envListenPort string = "LISTEN_PORT"

var config = &Config{}

func init() {
	config.ListenPort = os.Getenv(envListenPort)
}

func GetConfig() *Config {
	return config
}
