package basic

// CommonListResult returned result
type CommonListResult struct {
	Total   int         `json:"total"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
}
