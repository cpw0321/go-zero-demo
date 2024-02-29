// Code generated by goctl. DO NOT EDIT.
package types

type BasicRes struct {
	RetCode int64       `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
	Data    interface{} `json:"data"`
}

type ListBasicRes struct {
	RetCode  int64       `json:"retCode"`
	RetMsg   string      `json:"retMsg"`
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	PageNum  int64       `json:"pageNum"`
	PageSize int64       `json:"pageSize"`
}

type ListInfosReq struct {
	PageNum  int    `form:"pageNum,optional"`
	PageSize int    `form:"pageSize,optional"`
	Sort     string `form:"sort,optional"` // asc，desc
	Field    string `form:"field,optional"`
}
