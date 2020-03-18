package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"hellas/common/constant"
	"hellas/common/setting"
	"hellas/dtos/common"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type Claims struct {
	AccountId string `json:"accountId"`
	jwt.StandardClaims
}

// 生成Token
func GenerateToken(accountId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	var claims Claims
	claims.AccountId = accountId
	claims.ExpiresAt = expireTime.Unix()
	claims.Issuer = "hellas"

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(setting.JwtSecret))

	return token, err
}

// 解析token
func AuthorityCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		freeAuth := false
		for _, path := range setting.FreeAuthority {
			if path == c.FullPath() {
				freeAuth = true
			}
		}
		// 不需要认证，直接进入Controller
		if freeAuth {
			c.Next()
			return
		}

		var baseResult common.BaseResult
		token := c.Request.Header.Get("Authorization")
		// token为空
		if token == "" {
			baseResult.ErrorDto.Errors = append(baseResult.ErrorDto.Errors, JoinMessages("","noAuth"))
			// 返回status401
			c.JSON(http.StatusUnauthorized, baseResult)
			c.Abort()
			return
		}

		// 解析token
		_, err :=JwtParse(token)
		// 解析失败
		if err != nil {
			baseResult.ErrorDto.Errors = append(baseResult.ErrorDto.Errors, JoinMessages("","noAuth"))
			// 返回status401
			c.JSON(http.StatusUnauthorized, baseResult)
			c.Abort()
			return
		}
		//else if time.Now().Unix() > tokenInfo.ExpiresAt {
		//	// token过期
		//	baseResult.ErrorDto.Errors = append(baseResult.ErrorDto.Errors, JoinMessages("","authTimeOut"))
		//	// 返回status401
		//	c.JSON(http.StatusUnauthorized, baseResult)
		//	c.Abort()
		//	return
		//}

		// 验证通过，进入Controller
		c.Next()
	}
}

// 解析token
func JwtParse(token string) (*Claims, error) {
	tokenInfo , err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(setting.JwtSecret),nil
	})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return tokenInfo.Claims.(*Claims), nil
}

// 解析token
func JwtParseUser(token string) (string, error) {
	tokenInfo , err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(setting.JwtSecret),nil
	})
	if err != nil {
		log.Print(err)
		return "", err
	}
	finToken := tokenInfo.Claims.(jwt.MapClaims)
	return finToken["accountId"].(string), err
}

// 生成随机盐
func CreateSalt() string {
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	salt := make([]rune, 50)
	for i := range salt {
		//b := new(big.Int).SetInt64(int64(len(defaultLetters)))
		num, _ := rand.Int(rand.Reader, big.NewInt(62))
		salt[i] = defaultLetters[num.Int64()]
	}

	return string(salt)
}

// 生成随机验证码
func CreateCaptchaCode() string {
	var defaultLetters = []rune("0123456789")
	captchaCode := make([]rune, 6)
	for i := range captchaCode {
		num, _ := rand.Int(rand.Reader, big.NewInt(10))
		captchaCode[i] = defaultLetters[num.Int64()]
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
