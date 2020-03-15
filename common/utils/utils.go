package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-playground/validator"
	"hellas/common/constant"
	"log"
	"math/rand"
	"strings"
)

// 生成随机盐
func CreateSalt() string {
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	salt := make([]rune, 50)
	for i := range salt {
		salt[i] = defaultLetters[rand.Intn(len(defaultLetters))]
	}

	return string(salt)
}

// 生成随机验证码
func CreateCaptchaCode() string {
	var defaultLetters = []rune("0123456789")
	captchaCode := make([]rune, 6)
	for i := range captchaCode {
		captchaCode[i] = defaultLetters[rand.Intn(len(defaultLetters))]
	}

	return string(captchaCode)
}

// 生成混淆密码（MD5）
func CreateMd5Password(salt string, password string) string {
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(salt))
	confusePw := m5.Sum(nil)
	return hex.EncodeToString(confusePw)
}

// 生成错误信息
func CreateMessages(fieldErrors validator.ValidationErrors) []string {
	var messages []string
	for _, err := range fieldErrors {
		messages = append(messages, JoinMessages(err.Field(), err.Tag()))
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
func JoinMessages(filed string, validate string) string {
	var msg string
	var name string
	// 找不到错误信息用物理名
	if constant.Messages[validate] == "" {
		msg = "{0}:" + validate
	} else {
		msg = constant.Messages[validate]
	}

	// 找不到项目名用物理名
	if constant.Fields[filed] == "" {
		name = filed
	} else {
		name = constant.Fields[filed]
	}

	return strings.Replace(msg, "{0}", name, -1)
}