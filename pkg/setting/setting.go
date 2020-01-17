// @Time    : 2020/1/17 23:37
// @Author  : smalldoraemon@qq.com
// @File    : setting.go

package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini : %v'", err)

	}
}
