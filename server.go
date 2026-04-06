package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"git.oriondev.fr/orion/status/config"
	"git.oriondev.fr/orion/status/services"
)

type PageData struct {
	Config config.Config
	Duration time.Duration
	Message []string
}

func getPage(
	w http.ResponseWriter,
	config config.Config,
	t *template.Template,
) {
	var message []string
	if _, err := os.Stat("message.txt"); err == nil {
		d, err := os.ReadFile("message.txt")
		if err != nil {
			log.Println("Can't open message.txt")
		} else {
			s := string(d)
			message := strings.Split(s, "\n")
			if len(message) == 0 {
				message = nil
			}
		}
	}

	t.Execute(w, PageData{
		Config: config,
		Duration: time.Since(config.StartTime),
		Message: message,
	})
}

func Run(conf config.Config) {
	t := renderTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		getPage(w, conf, t)
	})

	conf.StartTime = time.Now()
	services.StartTimer(&conf.Services, conf.Interval)

	log.Printf("Listening on the port %d\n", conf.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
