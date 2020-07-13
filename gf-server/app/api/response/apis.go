package response

import "gf-server/app/model/apis"

type ApiListResponse struct {
	Apis []*apis.Entity `json:"apis"`
}

type ApiResponse struct {
	Api *apis.Entity `json:"api"`
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}