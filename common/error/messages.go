package error

import (
	"github.com/go-playground/validator"
	"log"
	"strings"
)

// 生成错误信息
func CreateMessages(fieldErrors validator.ValidationErrors) []string {
	var messages []string
	for _, err := range fieldErrors {
		messages = append(messages, joinMessages(err.Field(), err.Tag()))
		log.Printf(err.Namespace())
		log.Printf(err.Field())
		log.Printf(err.StructNamespace())
		log.Printf(err.StructField())
		log.Printf(err.Tag())
		log.Printf(err.ActualTag())
	}
	return messages
}

// 错误信息拼接
func joinMessages(filed string, validate string) string {
	var msg string
	var name string
	// 找不到错误信息
	if Messages[validate] == "" {
		msg = "{0}:" + validate
	} else {
		msg = Messages[validate]
	}

	// 找不到项目名用物理名
	if Fields[filed] == "" {
		name = filed
	} else {
		name = Fields[filed]
	}

	return strings.Replace(msg, "{0}", name, -1)
}


