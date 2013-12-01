package o2o

import (
	"github.com/zhhigh/o2o/config"
	"os"
	"path"
	"runtime"
	"fmt"
)
var (
	//AppName       string
	AppPath       string
	AppConfigPath string
	//Log conf
	LogFileName   string
	LogType       string
	//weixin conf,read from conf
	WX_DB        string
	WX_Collection string
	WX_Document   string
	WX_Token      string
	WX_Port      string
	WX_Defaultreply  string

	AppConfig     config.ConfigContainer

)

func init() {
	os.Chdir(path.Dir(os.Args[0]))
	AppPath = path.Dir(os.Args[0])
	AppConfigPath = path.Join(AppPath, "conf", "app.conf")
	err := ParseConfig()
	if (err != nil){
       fmt.Println("parse config error:",err)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

/*
logtype=file
logfile=test.log
jobserver= 106.187.94.189:14079
wx_db= mongodb://redis:redis@192.241.225.170:27017/rrest
wx_collection= rrest
wx_document= msg
wx_token= qingzhonglele
wx_port= :80
wx_defaultreply= 欢迎关注！订阅图文信息输入日期，如0901
*/

func ParseConfig() (err error) {
	AppConfig, err = config.NewConfig("ini", AppConfigPath)
	if err != nil {
		return err
	} else {
		if v := AppConfig.String("logfile"); err == nil {
			LogFileName = v
		}
		if v := AppConfig.String("logtype"); err == nil {
			LogType = v
		}
        if v := AppConfig.String("wx_db"); err == nil {
			WX_DB = v
		}
		if v := AppConfig.String("wx_collection"); err == nil {
			WX_Collection = v
		}
		if v := AppConfig.String("wx_document"); err == nil {
			WX_Document = v
		}
		if v := AppConfig.String("wx_token"); err == nil {
			WX_Token = v
		}
		if v := AppConfig.String("wx_port"); err == nil {
			WX_Port = v
		}
		if v := AppConfig.String("wx_defaultreply"); err == nil {
			WX_Defaultreply = v
		}
	}
	return nil
}
