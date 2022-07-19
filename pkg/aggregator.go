package pkg

import (
	"fmt"
	"strconv"
	"time"
)

func AggregateData() {
	PersonUsage = make(map[int64][]Usage)
	for {
		fmt.Println("Aggregating...")
		usages := <-UsagesChannel
		AppendUsages(usages)
	}
}

func AppendUsages(usages []Usage) {
	for j := 0; j < len(usages); j++ {
		PersonUsage[usages[j].Uid] = append(PersonUsage[usages[j].Uid], usages[j])
	}
}

var PersonUsage map[int64][]Usage

func PrintConsumers() {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
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
	}
}

func emptyUsages() {
	PersonUsage = make(map[int64][]Usage)
}
