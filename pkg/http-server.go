package pkg

import (
	"net/http"
)

const Port = ":8000"

type ErrorHandler interface {
	Error() string
	StatusCode() int
	APIError() (int, string)
}

type sentinelAPIError struct {
	status int
	msg    string
}

var _ ErrorHandler = sentinelAPIError{}

func (e sentinelAPIError) Error() string {
	return e.msg
}

func (e sentinelAPIError) StatusCode() int {
	return e.status
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.status, e.msg
}

var (
	ErrAuth = &sentinelAPIError{status: http.StatusUnauthorized, msg: "invalid token"}
)

func InitialServer() {
	InitialServices()
	RunAPI()
	ListenAndServe()
	wg.Wait()
}

func RunAPI() {
	Usages()
	Costs()
}

func ListenAndServe() {
	err := http.ListenAndServe(Port, nil)
	check(err)
}
