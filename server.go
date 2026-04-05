package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"git.oriondev.fr/orion/status/config"
	"git.oriondev.fr/orion/status/services"
)

func getPage(
	w http.ResponseWriter,
	config config.Config,
	t *template.Template,
) {
	if _, err := os.Stat("message.txt"); err == nil {
		d, err := os.ReadFile("message.txt")
		if err != nil {
			log.Println("Can't open message.txt")
		} else {
			s := string(d)
			msg := strings.Split(s, "\n")
			if len(msg) == 0 {
				msg = nil
			}
			config.Message = msg
		}
	}

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
