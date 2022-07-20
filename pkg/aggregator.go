package pkg

import (
	"fmt"
	"strconv"
)

//type CostsResponse struct {
//	PerService map[string]int64 `json:"per_service"`
//	Total      int64            `json:"total"`
//}

type UsagesResponse struct {
	Usages []Usage `json:"usages"`
}

var PersonUsage map[int64][]Usage
var Coefficients map[string]map[string]int64
var firstInitial = true

func AggregateData() {
	if firstInitial {
		emptyUsages()
	}
	for {
		newUsages := <-UsagesChannel
		AppendUsages(newUsages)
	}
}

func AppendUsages(usages []Usage) {
	for j := 0; j < len(usages); j++ {
		PersonUsage[usages[j].Uid] = append(PersonUsage[usages[j].Uid], usages[j])
	}
}

func CalculateConsumerUsages(uid int64) []Usage {
	usages := PersonUsage[uid]
	return usages
}

func CalculateConsumerCosts(uid int64) (map[string]int64, int64) {
	CostsPerService := make(map[string]int64)
	for _, usage := range PersonUsage[uid] {
		for tagName, value := range usage.Tags {
			CostsPerService[usage.Service] += value * (Coefficients[usage.Service][tagName])
		}
	}
	var total int64
	for _, value := range CostsPerService {
		total += value
	}
	return CostsPerService, total
}

func PrintConsumers() {
	for i, usages := range PersonUsage {
		var out string
		out += "U_ID: " + strconv.FormatInt(i, 10) + "\nUsages:\n"
		for _, usage := range usages {
			out += usage.Service + " - "
		}
		fmt.Println(out, "\n")
	}
}

func emptyUsages() {
	firstInitial = false
	PersonUsage = make(map[int64][]Usage)
}
