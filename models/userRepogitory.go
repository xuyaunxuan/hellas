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
	// 用户创捷返回结果
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
	// 用户创捷返回结果
	var result common.BaseResult
	// 用户信息DTO
	var userList []dtos.User
	// 邮箱地址存在验证
	if param.MailAddress != "" {
		db.Where(&dtos.User{MailAddress: param.MailAddress}).Find(&userList)
		// 查询结果为0
		if len(userList) == 0 {
			// 邮箱地址已被使用
			result.Result = constant.NG
			result.Errors = append(result.Errors, utils.JoinMessages("MailAddress","unused"))
			return result
		}
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

	return result
}