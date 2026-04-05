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
		return conf, errors.New("Configuration file not found. Please create a config.toml file")
	}

	_, err := toml.DecodeFile(f, &conf)
	if err != nil {
		return conf, err
	}

	if conf.Icon == "" {
		conf.Icon = "https://placehold.co/600x400"
	}

	if conf.Port == 0 {
		conf.Port = 3333
	}

	if conf.Interval == 0 {
		conf.Interval = 2
	}

	return conf, nil
}
