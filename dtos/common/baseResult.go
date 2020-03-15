package common

type BaseResult struct {
	ErrorDto `json:"errorDto"`
	Result string `json:"result"`
}