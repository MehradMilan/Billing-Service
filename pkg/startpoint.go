package pkg

import (
	"Billing/pkg/metric"
	"sync"
)

var wg sync.WaitGroup
var mux sync.Mutex

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func InitialServices() {
	UsagesChannel = make(chan []Usage, EndPointCount)
	InitialFileExtracts()
	go AggregateData()
	go CollectData(EndpointsAddress)
}

func ReportUsages(uid int64) UsagesResponse {
	consumerUsages := CalculateConsumerUsages(uid)
	usagesResponse.Usages = consumerUsages
	return usagesResponse
}

var costsResponse metric.CostsResponse
var usagesResponse UsagesResponse

func ReportCosts(uid int64) metric.CostsResponse {
	costsPerService, total := CalculateConsumerCosts(uid)
	costsResponse.PerService = costsPerService
	costsResponse.Total = total
	return costsResponse
}
