package pkg

import (
	"net/http"
)

const Port = ":8000"

func InitialServer() {
	RunAPI()
	ListenAndServe()
}

func RunAPI() {
	Usages()
	Costs()
}

func ListenAndServe() {
	err := http.ListenAndServe(Port, nil)
	check(err)
}
