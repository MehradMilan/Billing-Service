package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Usage struct {
	Service   string         `json:"service"`
	Tags      map[string]int `json:"tags"`
	Uid       int64          `json:"uid"`
	Timestamp int64          `json:"timestamp"`
}

var wg sync.WaitGroup
var mux sync.Mutex

type Endpoint struct {
	URLs []string
}

var EndPointCount int
var ServicesChannel chan []Usage

func ExtractEndpointsFromFile(address string) Endpoint {
	jsonFile, err := os.Open(address)
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var endpoint Endpoint
	err = json.Unmarshal(byteValue, &endpoint)
	check(err)
	return endpoint
}

func CollectData(address string) {
	endpoint := ExtractEndpointsFromFile(address)
	urls := endpoint.URLs
	EndPointCount = len(urls)

	ServicesChannel = make(chan []Usage, EndPointCount)

	wg.Add(EndPointCount)
	for i := 0; i < EndPointCount; i++ {
		go ProcessRequests(urls[i])
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
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)
	for {
		select {
		case <-done:
			wg.Done()
			return
		case <-ticker.C:
			mux.Lock()
			response := SendRequest(URL)
			services := DecodeData(response)
			ServicesChannel <- services
			mux.Unlock()
		}
	}
}

func DecodeData(jsonString string) []Usage {
	var services []Usage
	err := json.Unmarshal([]byte(jsonString), &services)
	check(err)
	fmt.Println(services)
	return services
}
