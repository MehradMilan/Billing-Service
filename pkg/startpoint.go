package pkg

import (
	"sync"
)

const EndpointsAddress = "./resources/endpoints.json"

var wg sync.WaitGroup
var mux sync.Mutex

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Start() {
	CollectData(EndpointsAddress)
	go AggregateData()
	wg.Add(EndPointCount)
	for i := 0; i < EndPointCount; i++ {
		go ProcessRequests(URLs[i])
	}
	wg.Wait()
}
