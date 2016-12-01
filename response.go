package models

type Response struct {
	Code int         `json:"httpstatus"`
	Data interface{} `json:"data"`
}
