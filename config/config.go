package config

import "github.com/pteich/configstruct"

type Config struct {
	LogLevel    string `cli:"logLevel" env:"LOG_LEVEL"`
	LogConsole  bool   `cli:"logConsole" env:"LOG_CONSOLE"`
	Pin         string `cli:"pin" env:"HOMEKIT_PIN"`
	ConfigFile  string `cli:"config" env:"CONFIG_FILE"`
	StoragePath string `cli:"storage" env:"STORAGE_PATH"`
}

func New() Config {
	cfg := Config{
		LogLevel:    "debug",
		LogConsole:  false,
		Pin:         "12345678",
		ConfigFile:  "./accessories.json",
		StoragePath: "./",
	}

	configstruct.Parse(&cfg)

	return cfg
}
