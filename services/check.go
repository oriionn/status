package services

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/tcnksm/go-httpstat"
)

func Fetch(method string, url string, result *httpstat.Result) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	ctx := httpstat.WithHTTPStat(req.Context(), result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	return client.Do(req)

}

func Check(service *Service) bool {
	log.Printf("Service: checking status of %s\n", service.Name)
	service.Total++

	var result httpstat.Result

	url := service.URL
	res, err := Fetch("HEAD", url, &result)
	if err != nil {
		res, err = Fetch("GET", url, &result)
		if err != nil {
			return false
		}
		io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}

	status := res.StatusCode < 400
	if status {
		service.Up++
	}

	service.Status = status
	service.Latency = int(result.StartTransfer / time.Millisecond)
	return status
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
