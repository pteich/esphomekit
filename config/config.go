package config

import "github.com/pteich/configstruct"

type Config struct {
	LogLevel   string `cli:"logLevel"`
	LogConsole bool   `cli:"logConsole"`
	Pin        string `cli:"pin"`
	ConfigFile string `cli:"config"`
}

func New() Config {
	cfg := Config{
		LogLevel:   "debug",
		LogConsole: false,
		Pin:        "12345678",
	}

	configstruct.Parse(&cfg)

	return cfg
}
