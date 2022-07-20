package pkg

import (
	"fmt"
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
