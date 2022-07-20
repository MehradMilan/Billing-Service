package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Usages() {
	http.HandleFunc("/usages", func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("uid")
		fmt.Println("id =>", uid)
		uidInt, err := strconv.Atoi(uid)
		check(err)
		response := ReportUsages(int64(uidInt))
		err = json.NewEncoder(w).Encode(response)
		check(err)
	})
}

func Costs() {
	http.HandleFunc("/costs", func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("uid")
		uidInt, err := strconv.Atoi(uid)
		check(err)
		response := ReportCosts(int64(uidInt))
		err = json.NewEncoder(w).Encode(response)
		check(err)
	})
}
