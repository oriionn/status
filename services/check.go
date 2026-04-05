package services

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func Check(service *Service) bool {
	log.Printf("Service: checking status of %s\n", service.Name)
	service.Total++

	url := service.URL
	res, err := http.Head(url)
	if err != nil {
		res, err = http.Get(url)
		if err != nil {
			return false
		}
		io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}

	up := res.StatusCode < 400
	if up {
		service.Up++
	}
	service.Status = up
	return up
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
