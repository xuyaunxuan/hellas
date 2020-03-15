package models

import (
	"bytes"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"hellas/common/constant"
	"hellas/common/setting"
	"hellas/common/utils"
	"hellas/dtos"
	"hellas/dtos/common"
	"hellas/dtos/user"
	"html/template"
	"time"
)

// 新用户创建
func CreateNewUser(param user.RegisterParameter) common.BaseResult {
	// 开启事务
	tx := db.Begin()
	// 用户创建返回结果
	var result common.BaseResult
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
	if param.MailAddress != "" {
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

// 发送用户验证邮件
func SendCaptchaMail(param user.SendCaptchaMailParameter) common.BaseResult {
	// 开启事务
	tx := db.Begin()
	// 用户验证返回结果
	var result common.BaseResult
	// 用户信息DTO
	var userList []dtos.User
	// 邮箱地址存在验证
	db.Where(&dtos.User{MailAddress: param.MailAddress}).Find(&userList)
	// 查询结果为0
	if len(userList) == 0 {
		// 邮箱地址不存在
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("MailAddress","unused"))
		return result
	}
	// 生成随机验证码
	var captchaCode = utils.CreateCaptchaCode()
	// 存储验证码
	db.Model(dtos.User{}).Where(dtos.User{MailAddress:param.MailAddress}).Update(dtos.User{CaptchaCode : captchaCode, SendCaptchaTime:time.Now()})
	// 提交事务
	tx.Commit()
	// 处理成功
	result.Result = constant.OK

	// 邮件发送
	m := gomail.NewMessage()
	// ikenLab <noreply@ikenlab.com>
	m.SetHeader("From", setting.MailMan + "<" +  setting.MailAddress + ">")
	m.SetHeader("To", param.MailAddress)
	// 标题
	m.SetHeader("Subject", "ikenLab - 验证邮件")
	// 内容模板生成
	t, _ := template.New("MAIL").ParseFiles("./common/template/captchaTemplate.html")
	// 读取模板
	buffer := new(bytes.Buffer)
	code := gin.H{
		"captchaCode": captchaCode,
		"nickName": userList[0].NickName,
	}
	// 验证码替换
	t.ExecuteTemplate(buffer, "captchaTemplate.html", code)
	// 邮件内容设定
	m.SetBody("text/html", buffer.String())
	// 邮件发送
	d := gomail.NewDialer(setting.MailHost, setting.MailPort, setting.MailAddress, setting.MailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	d.DialAndSend(m)
	// 处理成功
	result.Result = constant.OK
	return result
}

// 重置密码
func ResetPassword(param user.ResetPasswordParameter) common.BaseResult {
	// 开启事务
	tx := db.Begin()
	// 用户验证返回结果
	var result common.BaseResult
	// 用户信息DTO
	var userList []dtos.User
	// 密码输入一致性验证
	if param.OncePassword != param.TwicePassword {
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","samePassword"))
		return result
	}
	// 邮箱地址存在验证
	db.Where(&dtos.User{MailAddress: param.MailAddress}).Find(&userList)
	// 查询结果为0
	if len(userList) == 0 {
		// 邮箱地址不存在
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("MailAddress","unused"))
		return result
	}
	// 验证码一致性验证
	if userList[0].CaptchaCode != param.CaptchaCode {
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","sameCaptcha"))
		return result
	}

	// 生成随机盐
	var salt = utils.CreateSalt()
	// 生成混淆密码
	var confusePassword = utils.CreateMd5Password(salt, param.OncePassword)

	// 更改密码，更改盐值
	db.Model(dtos.User{}).Where(dtos.User{MailAddress:param.MailAddress}).Update(dtos.User{Salt : salt, Password:confusePassword, UpdateDateTime:time.Now(),CaptchaCode:"######"})

	// 提交事务
	tx.Commit()
	// 处理成功
	result.Result = constant.OK
	return result
}

// 用户登录
func Login(param user.LoginParameter) user.LoginResult {
	// 开启事务
	tx := db.Begin()
	// 用户登录返回结果
	var result user.LoginResult
	// 用户信息DTO
	var userList []dtos.User
	// ID/邮箱必须输入一个
	if param.AccountId == "" && param.MailAddress == "" {
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","idOrMailMust"))
		return result
	}

	// ID登录
	if param.AccountId != "" {
		db.Where(&dtos.User{AccountId: param.AccountId}).Find(&userList)
	} else if param.MailAddress != "" {
		db.Where(&dtos.User{MailAddress: param.MailAddress}).Find(&userList)
	}

	// ID/邮箱存在验证
	if len(userList) == 0 {
		// ID/邮箱不存在
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","idOrMailNotExist"))
		return result
	}

	// 生成混淆密码
	var confusePassword = utils.CreateMd5Password(userList[0].Salt, param.Password)
	// 密码验证
	if confusePassword != userList[0].Password {
		// 密码不正确
		result.Result = constant.NG
		result.Errors = append(result.Errors, utils.JoinMessages("","passwordIncorrect"))
		return result
	}

	// 用户ID设定
	result.AccountId = userList[0].AccountId
	// 用户邮箱设定
	result.MailAddress = userList[0].MailAddress
	// 用户昵称设定
	result.NickName = userList[0].NickName
	// token生成
	token, _ := utils.GenerateToken(userList[0].AccountId)
	// token设定
	result.Token = token
	// 提交事务
	tx.Commit()
	// 处理成功
	result.Result = constant.OK
	return result
}