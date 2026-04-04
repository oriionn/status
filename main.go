package main

import (
	_ "embed"
	"html/template"
	"net/http"
	"time"

	"git.oriondev.fr/orion/status/services"
)

type PageData struct {
	Services []services.Service
}



func getPage(
	w http.ResponseWriter,
	servicesList []services.Service,
	t *template.Template,
) {
	t.Execute(w, PageData{
		Services: servicesList,
	})
}

func main() {
	servicesList := []services.Service{
		{
			Name: "Forgejo",
			URL: "https://git.oriondev.fr",
			ShowURL: true,
		},
		{
			Name: "Portfolio",
			URL: "https://oriondev.f",
			ShowURL: true,
		},
	}

	t := renderTemplate()

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		getPage(w, servicesList, t)
	})

	go func() {
		for range time.Tick(time.Second) {
			services.CheckServices(&servicesList)
		}
	}()

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
