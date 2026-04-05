package main

import (
	_ "embed"
	"html/template"
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

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	t := renderTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		getPage(w, conf, t)
	})

	services.StartTimer(&conf.Services, conf.Interval)

	err = http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
