package common

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Config       *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JWTSecret    string
)

func init() {
	var err error
	Config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to load config: %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Config.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Config.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to load server: %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Config.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to load app: %v", err)
	}

	JWTSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
