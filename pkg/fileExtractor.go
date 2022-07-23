package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ExtractCoefficients(address string) {
	jsonFile, err := os.Open(address)
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Coefficients)
	fmt.Println(Coefficients)
	check(err)
}

func ExtractEndpointsFromFile(address string) Endpoint {
	jsonFile, err := os.Open(address)
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var endpoint Endpoint
	err = json.Unmarshal(byteValue, &endpoint)
	check(err)
	return endpoint
}
