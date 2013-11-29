package o2o

import (
	"github.com/zhhigh/o2o/config"
	"os"
	"path"
	"runtime"
)

var (
	//AppName       string
	AppPath       string
	AppConfigPath string
	HttpPort      string
	LogFileName   string
	LogType       string
	AppConfig     config.ConfigContainer

)

func init() {
	os.Chdir(path.Dir(os.Args[0]))
	AppPath = path.Dir(os.Args[0])
	AppConfigPath = path.Join(AppPath, "conf", "app.conf")
	ParseConfig()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func ParseConfig() (err error) {
	AppConfig, err = config.NewConfig("ini", AppConfigPath)
	if err != nil {
		return err
	} else {
		if v := AppConfig.String("httpport"); err == nil {
			HttpPort = v
		}
		if v := AppConfig.String("logfile"); err == nil {
			LogFileName = v
		}
		if v := AppConfig.String("logtype"); err == nil {
			LogType = v
		}

	}
	return nil
}
