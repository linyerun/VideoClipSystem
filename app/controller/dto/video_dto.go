package dto

// swagger:model VideoDto
type VideoDto struct {
	Url       string `json:"url"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}
