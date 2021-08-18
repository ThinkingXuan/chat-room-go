package rr

type ReqPage struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type ReqClusterIP []string
