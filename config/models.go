package config

import "git.oriondev.fr/orion/status/services"

type Config struct {
	Title string `toml:"title"`
	Description string `toml:"description"`
	Icon string `toml:"icon"`
	Interval int `toml:"interval"`
	Services []services.Service `toml:"service"`
}
