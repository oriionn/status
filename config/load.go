package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

func Load() (Config, error) {
	var conf Config
	f := "config.toml"
	if _, err := os.Stat("config.toml"); err != nil {
		return conf, errors.New("Configuration file not file. Please create a config.toml file")
	}

	_, err := toml.DecodeFile(f, &conf)
	return conf, err
}
