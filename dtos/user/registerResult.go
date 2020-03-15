package user

import (
	"hellas/dtos/common"
)
type RegisterResult struct {
	common.ErrorDto
	Result string
}