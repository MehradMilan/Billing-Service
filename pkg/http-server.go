package pkg

import (
	"fmt"
	"net/http"
	"strconv"
)

const Port = 8000

func InitialServer() {
	RunAPI()
	ListenAndServe()
}

func RunAPI() {
	http.HandleFunc("/usages", func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("uid")
		fmt.Println("id =>", uid)
		uidInt, err := strconv.Atoi(uid)
		check(err)
		ReportUsages(int64(uidInt))
	})
}

func ListenAndServe() {
	err := http.ListenAndServe(":8000", nil)
	check(err)
}
