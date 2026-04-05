package main

import (
	_ "embed"
	"log"

	"git.oriondev.fr/orion/status/config"
	flag "github.com/spf13/pflag"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	flag.IntVarP(&conf.Port, "port", "p", conf.Port, "listening port")
	flag.IntVarP(&conf.Interval, "interval", "i", conf.Interval, "checking interval (in seconds)")
	flag.Parse()

	Run(conf)
}
