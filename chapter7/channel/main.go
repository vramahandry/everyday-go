package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var uri string
	flag.StringVar(&uri, "uri", "", "uri to fetch")
	flag.Parse()
	titleCh := make(chan string)
	errorCh := make(chan error)
	go titleOf(uri, titleCh, errorCh)
	select {
	case title := <-titleCh:
		fmt.Printf("Title: %s\n", title)
	case err := <-errorCh:
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func titleOf(uri string, titleCh chan string, errCh chan error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		errCh <- err
		return
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		errCh <- err
		return
	}
	if res.Body != nil {
		defer res.Body.Close()
		res, _ := io.ReadAll(res.Body)
		for _, l := range strings.Split(string(res), "\n") {
			if strings.Contains(l, "<title>") {
				titleCh <- l
				return
			}
		}
	}
	errCh <- fmt.Errorf("no title found")
}
