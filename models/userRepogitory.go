package models

import (
	"hellas/common/constant"
	"hellas/common/utils"
	"hellas/dtos"
	"hellas/dtos/user"
	"log"
	"time"
)


type Login struct {
	Id string
}

type Product struct {
	Code string
	Price uint
}

// 新用户创建
func CreateNewUser(param user.RegisterParameter) user.RegisterResult {
	// 开启事务
	tx := db.Begin()
	// 用户创捷返回结果
	var result user.RegisterResult
	// 用户信息DTO
	var userList []dtos.User
	// 用户ID唯一性验证
	if param.AccountId != "" {
		db.Where(&dtos.User{AccountId: param.AccountId}).Find(&userList)
		// 查询结果大于0
		if len(userList) > 0 {
			// 账户ID已被使用
			result.Result = constant.NG
			result.Errors = append(result.Errors, utils.JoinMessages("AccountId","used"))
			return result
		}
	}

	// 邮箱地址唯一性验证
	if param.AccountId != "" {
		db.Where(&dtos.User{MailAddress: param.MailAddress}).Find(&userList)
		// 查询结果大于0
		if len(userList) > 0 {
			// 邮箱地址已被使用
			result.Result = constant.NG
			result.Errors = append(result.Errors, utils.JoinMessages("MailAddress","used"))
			return result
		}
	}
	// 生成随机盐
	var salt = utils.CreateSalt()
	// 生成混淆密码
	var confusePassword = utils.CreateMd5Password(salt, param.Password)
	var user = dtos.User{
		AccountId: param.AccountId,
		MailAddress: param.MailAddress,
		NickName: param.NickName,
		Password: confusePassword,
		Salt: salt,
		InsertDateTime:time.Now(),
		UpdateDateTime: time.Now(),
	}

	// 用户创建
	db.Create(&user)
	// 提交事务
	tx.Commit()
	// 处理成功
	result.Result = constant.OK
	return result
}


func GetUser() []Product {
	var login []Login

	db.Debug().Find(&login)
	log.Printf("%+v", login)
	//
	//db.AutoMigrate(&Product{})
	//
	//db.Create(&Product{Code: "L1212", Price: 1000})
	var product []Product
	db.Find(&product) // find product with id 1
	log.Printf("%+v", product)
	return product
}