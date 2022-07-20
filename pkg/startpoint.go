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

func ReportUsages(uid int64) {
	CollectData(EndpointsAddress)
	AggregateData()
	PrintSelectedConsumerUsages(uid)
}
