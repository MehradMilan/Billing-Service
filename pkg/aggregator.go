package pkg

import (
	"fmt"
	"strconv"
)

var AllUsages [][]Usage
var flag bool

func AggregateData() {
	for {
		for len(AllUsages) < EndPointCount {
			flag = true
			fmt.Println(len(AllUsages))
			AllUsages = append(AllUsages, <-UsagesChannel)
		}
		if flag {
			PrintConsumers()
		}
	}
}

var PersonUsage map[int64][]Usage

func PrintConsumers() {
	flag = false
	PersonUsage = make(map[int64][]Usage)
	for i := 0; i < len(AllUsages); i++ {
		for j := 0; j < len(AllUsages[i]); j++ {
			PersonUsage[AllUsages[i][j].Uid] = append(PersonUsage[AllUsages[i][j].Uid], AllUsages[i][j])
		}
	}
	for i, usages := range PersonUsage {
		var out string
		out += "U_ID: " + strconv.FormatInt(i, 10) + "\nUsages:\n"
		for _, usage := range usages {
			out += usage.Service + " - "
		}
		fmt.Println(out, "\n")
	}
}
