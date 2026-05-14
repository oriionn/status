package fetch

import (
	"io"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

func fetch(method string, url string, result *httpstat.Result) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	ctx := httpstat.WithHTTPStat(req.Context(), result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	return client.Do(req)
}

func Http(url string) (Data, error) {
	var result httpstat.Result

	res, err := fetch("HEAD", url, &result)
	if err != nil {
		res, err = fetch("GET", url, &result)
		if err != nil {
			return Data{}, err
		}
		io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}

	status := res.StatusCode < 400
	latency := int(result.StartTransfer / time.Millisecond)

	return Data{
		Status: status,
		Latency: latency,
	}, nil
}
