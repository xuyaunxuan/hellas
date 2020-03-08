package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type BaseController struct {
	beego.Controller
}

type ControllerError struct {
	UrlLong  string
}

func (c *BaseController) ParseToken() (t *jwt.Token, e *ControllerError) {
	authString := c.Ctx.Input.Header("Authorization")
	var errInputData  ControllerError
	errInputData.UrlLong = "1"
	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		return nil, &errInputData
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, &errInputData
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, &errInputData
			} else {
				// Couldn't handle this token
				return nil, &errInputData
			}
		} else {
			// Couldn't handle this token
			return nil, &errInputData
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		return nil, &errInputData
	}
	beego.Debug("Token:", token)

	return token, nil
}