package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("User-Agent", "spacecount-tutorial")
	res, err := spaceClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	p := people{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Number)
}
