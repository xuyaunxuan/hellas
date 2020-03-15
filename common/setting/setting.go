package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	MailMan string
	MailAddress string
	MailPassword string
	MailHost string
	MailPort int

	DataBaseType string
	DataBaseName string
	DataBasePath string


	PageSize int
	JwtSecret string
)

func Init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	// 邮件设定
	LoadMail()
	// 数据库设定
	LoadDataBase()
	//LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 邮件设定
func LoadMail() {
	sec, err := Cfg.GetSection("mail")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	MailMan = sec.Key("MAIL_MAN").MustString("")
	MailAddress = sec.Key("MAIL_ADDRESS").MustString("")
	MailPassword = sec.Key("MAIL_PASSWORD").MustString("")
	MailHost = sec.Key("MAIL_HOST").MustString("")
	MailPort = sec.Key("MAIL_PORT").MustInt(465)
}

// 数据库设定
func LoadDataBase() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	DataBaseType = sec.Key("DATABASE_TYPE").MustString("")
	DataBaseName = sec.Key("DATABASE_NAME").MustString("")
	DataBasePath = sec.Key("DATABASE_PATH").MustString("")
}

//func LoadApp() {
//	sec, err := Cfg.GetSection("app")
//	if err != nil {
//		log.Fatalf("Fail to get section 'app': %v", err)
//	}
//
//	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
//	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
//}