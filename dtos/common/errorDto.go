package common

type ErrorDto struct {
	Message string  `json:"message"`
	Errors []string `json:"errors"`
}