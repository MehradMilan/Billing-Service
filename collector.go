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

// type Tags struct {
// 	tage_1 int
// 	tage_2 int
// 	tage_3 int
// }

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

// func readFileLineByLine(address string) []string {
// 	var fileLines []string
// 	content, err := os.Open(address)
// 	check(err)
// 	scanner := bufio.NewScanner(content)
// 	for scanner.Scan() {
// 		s := scanner.Text()
// 		fileLines = append(fileLines, s)
// 	}
// 	return fileLines
// }

func main() {
	endpoint := ExtractEndpointsFromFile("endpoints.json")
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
		// fmt.Println(out)
		DecodeData(out)
		check(scanner.Err())
		fmt.Println("\n\n")
		mux.Unlock()

		time.Sleep(5 * time.Second)
	}
	wg.Done()
}

func DecodeData(jsonString string) {
	// fmt.Println(jsonString)
	var services []Service
	// ss := []Service{}
	err := json.Unmarshal([]byte(jsonString), &services)
	check(err)
	fmt.Println(services)
}
