package main

import (
	"fmt"
	"io"

	"net/http"
	"sync"
)

type ScrapeRun struct {
	Results map[string]ScrapeResult
	Lock    *sync.Mutex
}
type ScrapeResult struct {
	HTTPCode   int
	HTTPLength int64
	Error      error
}

func scrape(url string, sr *ScrapeRun) {
	res, err := http.Get(url)
	if err != nil {
		sr.Lock.Lock()
		defer sr.Lock.Unlock()
		sr.Results[url] = ScrapeResult{
			Error:    err,
			HTTPCode: http.StatusBadGateway,
		}
		return
	}
	defer res.Body.Close()
	sr.Lock.Lock()
	defer sr.Lock.Unlock()
	length := res.ContentLength
	// Read whole body to find length, if not returned in
	// response
	if length == -1 {
		body, _ := io.ReadAll(res.Body)
		length = int64(len(body))
	}
	sr.Results[url] = ScrapeResult{
		HTTPCode:   res.StatusCode,
		HTTPLength: length,
	}
}

func main() {

	urls := []string{"https://inlets.dev/", "https://openfaas.com/", "https://openfas.com/"}
	sr := ScrapeRun{
		Lock:    &sync.Mutex{},
		Results: make(map[string]ScrapeResult),
	}
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		u := url
		go func(u string) {
			defer wg.Done()
			scrape(u, &sr)

		}(u)

	}
	wg.Wait()

	for k, v := range sr.Results {
		if v.Error != nil {
			fmt.Printf("- %s = error: %s\n", k, v.Error.Error())
		} else {
			fmt.Printf("- %s = [%d] %d bytes\n", k, v.HTTPCode, v.HTTPLength)
		}
	}
}
