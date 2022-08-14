package apipart

import (
	"encoding/json"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact string
}

func getCatFact(url string) (w string) {
	var catFact CatFact
	err := getJson(url, &catFact)
	if err != nil {
		w = "there was an error getting the fact: " + err.Error()
	} else {
		w = "the fact is: " + catFact.Fact
	}
	return

}
func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func Runner() string {
	url := "https://catfact.ninja/fact"
	client = &http.Client{Timeout: 10 * time.Second}
	return getCatFact(url)
}
