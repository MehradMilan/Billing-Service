package pkg

import (
	"sync"
)

const EndpointsAddress = "./resources/endpoints.json"
const ServicesCostsAddress = "./resources/coefficients.json"

var wg sync.WaitGroup
var mux sync.Mutex

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitialServices() {
	UsagesChannel = make(chan []Usage, EndPointCount)
	go AggregateData()
	go CollectData(EndpointsAddress)
}

func ReportUsages(uid int64) UsagesResponse {
	consumerUsages := CalculateConsumerUsages(uid)
	usagesResponse.Usages = consumerUsages
	return usagesResponse
}

var costsResponse CostsResponse
var usagesResponse UsagesResponse

func ReportCosts(uid int64) CostsResponse {
	costsPerService, total := CalculateConsumerCosts(uid)
	costsResponse.PerService = costsPerService
	costsResponse.Total = total
	return costsResponse
}
