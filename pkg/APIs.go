package pkg

import (
	"encoding/json"
	"net/http"
)

func Usages() {
	http.HandleFunc("/usages", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		uid := AuthenticateUser(token)
		if uid == -1 {
			response := http.Response{
				Status:     ErrAuth.Error(),
				StatusCode: ErrAuth.StatusCode(),
			}
			w.WriteHeader(ErrAuth.StatusCode())
			err := json.NewEncoder(w).Encode(response)
			check(err)
			return
		}
		response := ReportUsages(int64(uid))
		err := json.NewEncoder(w).Encode(response)
		check(err)
	})
}

func Costs() {
	http.HandleFunc("/costs", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		uid := AuthenticateUser(token)
		if uid == -1 {
			response := http.Response{
				Status:     ErrAuth.Error(),
				StatusCode: ErrAuth.StatusCode(),
			}
			err := json.NewEncoder(w).Encode(response)
			check(err)
			return
		}
		response := ReportCosts(int64(uid))
		err := json.NewEncoder(w).Encode(response)
		check(err)
	})
}
