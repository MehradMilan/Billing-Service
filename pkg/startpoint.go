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

func ReportUsages(uid int64) {
	CollectData(EndpointsAddress)
	AggregateData()
	PrintSelectedConsumerUsages(uid)
}

var response Response

func ReportCosts(uid int64) Response {
	CollectData(EndpointsAddress)
	AggregateData()
	costsPerService, total := CalculateConsumerCosts(uid)
	response.PerService = costsPerService
	response.Total = total
	return response
}
