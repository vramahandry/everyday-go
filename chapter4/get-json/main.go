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

func getAstros(apiURL string) (people, error) {

	p := people{}

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)

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

	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal(err)
	}
	return p, nil

}
func main() {

	apiURL := "http://api.open-notify.org/astros.json"
	people, err := getAstros(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d people found in space.\n", people.Number)

}
