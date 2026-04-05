package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"git.oriondev.fr/orion/status/config"
	"git.oriondev.fr/orion/status/services"
)

func getPage(
	w http.ResponseWriter,
	config config.Config,
	t *template.Template,
) {
	t.Execute(w, config)
}

func Run(conf config.Config) {
	t := renderTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		getPage(w, conf, t)
	})

	services.StartTimer(&conf.Services, conf.Interval)

	log.Printf("Listening on the port %d\n", conf.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
