package services

import (
	"log"
	"sync"
	"time"

	"git.oriondev.fr/orion/status/services/fetch"
)



func Check(service *Service) bool {
	log.Printf("Service: checking status of %s\n", service.Name)
	service.Total++

	data, err := fetch.Fetch(service.URL)
	if err != nil {
		return false
	}

	if data.Status {
		service.Up++
	}

	service.Status = data.Status
	service.Latency = data.Latency
	return data.Status
}

func CheckServices(l *[]Service) {
	var wg sync.WaitGroup
	for i := range *l {
		wg.Add(1)
		go func(s *Service) {
			defer wg.Done()
			a := Check(s)
			log.Printf("Service: %s's status = %t\n", s.Name, a)
		}(&(*l)[i])
	}
	wg.Wait()
}

func StartTimer(l *[]Service, interval int) {
	CheckServices(l)
	go func() {
		for range time.Tick(time.Duration(interval) * time.Second) {
			CheckServices(l)
		}
	}()
}
