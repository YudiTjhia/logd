package api

type ApiModel struct {
	Q string `json:"q"`
	UsePaging bool `json:"usePaging"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
}
