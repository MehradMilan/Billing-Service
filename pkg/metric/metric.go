package metric

type CostsResponse struct {
	PerService map[string]int64 `json:"per_service"`
	Total      int64            `json:"total"`
}
