package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const EndpointsAddress = "./resources/endpoints.json"
const ServicesCostsAddress = "./resources/coefficients.json"
const HashedTokensAddress = "./resources/auth-file.json"

func InitialFileExtracts() {
	ExtractHashedTokens(HashedTokensAddress)
	ExtractCoefficients(ServicesCostsAddress)
	ExtractEndpointsFromFile(EndpointsAddress)
}

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
	err = json.Unmarshal(byteValue, &endpoint)
	check(err)
	return endpoint
}

func ExtractHashedTokens(address string) {
	jsonFile, err := os.Open(address)
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &HashedTokens)
	check(err)
}
