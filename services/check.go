package services

import (
	"io"
	"net/http"
	"sync"
)

func Check(service *Service) bool {
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
			Check(s)
		}(&(*l)[i])
	}
	wg.Wait()
}
