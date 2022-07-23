package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Usage struct {
	Service   string           `json:"service"`
	Tags      map[string]int64 `json:"tags"`
	Uid       int64            `json:"uid"`
	Timestamp int64            `json:"timestamp"`
}

type Endpoint struct {
	URLs []string
}

const collectInterval = 5

var URLs []string
var EndPointCount int
var UsagesChannel chan []Usage

func CollectData(address string) {
	endpoint := ExtractEndpointsFromFile(address)
	URLs = endpoint.URLs
	EndPointCount = len(URLs)
	wg.Add(EndPointCount)
	for i := 0; i < EndPointCount; i++ {
		go ProcessRequests(URLs[i])
	}
	wg.Wait()
}

func SendRequest(URL string) string {
	resp, err := http.Get(URL)
	check(err)
	defer resp.Body.Close()

	fmt.Println("Response status:\n", resp.Status)
	var out string
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		out += scanner.Text()
	}
	check(scanner.Err())

	return out
}

func ProcessRequests(URL string) {
	ticker := time.NewTicker(collectInterval * time.Second)
	done := make(chan bool)
	for {
		select {
		case <-done:
			wg.Done()
			return
		case <-ticker.C:
			mux.Lock()
			response := SendRequest(URL)
			usages := DecodeResponse(response)
			UsagesChannel <- usages
			mux.Unlock()
		}
	}
}

func DecodeResponse(jsonString string) []Usage {
	var usages []Usage
	err := json.Unmarshal([]byte(jsonString), &usages)
	check(err)
	return usages
}
