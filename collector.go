package main

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

type Service struct {
	Name      string         `json:"service"`
	Tags      map[string]int `json:"tags"`
	Uid       int64          `json:"uid"`
	Timestamp int64          `json:"timestamp"`
}

var wg sync.WaitGroup
var mux sync.Mutex

type Endpoint struct {
	URLs []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExtractEndpointsFromFile(address string) Endpoint {
	jsonFile, err := os.Open(address)
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var endpoint Endpoint
	erro := json.Unmarshal(byteValue, &endpoint)
	check(erro)
	return endpoint
}

func main() {
	endpoint := ExtractEndpointsFromFile("./resources/endpoints.json")
	urls := endpoint.URLs

	wg.Add(len(urls))
	for i := 0; i < len(urls); i++ {
		go CollectData(urls[i])
	}
	wg.Wait()
}

func CollectData(URL string) {
	for i := 0; i < 100; i++ {
		resp, err := http.Get(URL)
		check(err)
		defer resp.Body.Close()

		mux.Lock()
		fmt.Println("Response status:\n", resp.Status)
		var out string
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan(); i++ {
			out += scanner.Text()
		}
		DecodeData(out)
		check(scanner.Err())
		fmt.Println()
		mux.Unlock()

		time.Sleep(5 * time.Second)
	}
	wg.Done()
}

func DecodeData(jsonString string) {
	var services []Service
	err := json.Unmarshal([]byte(jsonString), &services)
	check(err)
	fmt.Println(services)
}
