package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	currentUrl := "https://swapi.dev/api/people/"
	// AsyncRequests(currentUrl)
	// SyncRequests(currentUrl)
	fmt.Println("end program")
}

func SyncRequests(currentUrl string) {
	currentPage := 1

	for currentPage <= 9 {
		resp, err := HttpRequest("GET", currentUrl)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(currentUrl)
		fmt.Println(resp.Results)
		currentUrl = resp.Next
		currentPage++
	}

}

func AsyncRequests(url string) {
	var wg sync.WaitGroup

	wg.Add(8)
	currentPage := 1

	for currentPage <= 9 {
		urlWithpage := fmt.Sprintf("%s?page=%d", url, currentPage)

		go func(currentUrl string) {
			defer wg.Done()

			resp, err := HttpRequest("GET", currentUrl)
			if err != nil {
				fmt.Println("Error", err)
				return
			}

			fmt.Println(currentUrl)
			fmt.Println(resp.Results)
		}(urlWithpage)

		currentPage++
	}

	wg.Wait()
}

type Person struct {
	Name      string `json:"name"`
	BirthYear string `json:"birth_year"`
}

type StartWarsResponse struct {
	Next    string   `json:"next"`
	Results []Person `json:"results"`
}

func HttpRequest(method string, url string) (*StartWarsResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response StartWarsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("status invalido")
	}
	return &response, err
}
