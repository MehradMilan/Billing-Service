package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func AggregateData() {
	emptyUsages()
	for {
		select {
		case newUsages := <-UsagesChannel:
			AppendUsages(newUsages)
			fmt.Println("Aggregating...")
		default:
			return
		}
	}
}

func AppendUsages(usages []Usage) {
	for j := 0; j < len(usages); j++ {
		PersonUsage[usages[j].Uid] = append(PersonUsage[usages[j].Uid], usages[j])
	}
}

var PersonUsage map[int64][]Usage

func PrintSelectedConsumerUsages(uid int64) {
	usages := PersonUsage[uid]
	var out string
	out += "U_ID: " + strconv.FormatInt(uid, 10) + "\nUsages:\n"
	for _, usage := range usages {
		out += usage.Service + " - "
	}
	fmt.Println(out, "\n")
}

var Coefficients map[string]map[string]int64

type Response struct {
	PerService map[string]int64 `json:"per_service"`
	Total      int64            `json:"total"`
}

func PrintSelectedConsumerCosts(uid int64) {
	//usages := PersonUsage[uid]
	ExtractCoefficients("./resources/coefficients.json")
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
	res.PerService = CostsPerService
	res.Total = total
	fmt.Println(res)
}

var res Response

func ExtractCoefficients(address string) {
	jsonFile, err := os.Open(address)
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Coefficients)
	fmt.Println(Coefficients)
	check(err)
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
	emptyUsages()
}

func emptyUsages() {
	PersonUsage = make(map[int64][]Usage)
}
