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

)

func init() {
	os.Chdir(path.Dir(os.Args[0]))
	AppPath = path.Dir(os.Args[0])
	//HttpPort = ":80"
	//AppName = "beego"
	AppConfigPath = path.Join(AppPath, "conf", "app.conf")
	ParseConfig()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func ParseConfig() (err error) {
	AppConfig, err = config.NewConfig("ini", AppConfigPath)
	if err != nil {
		return err
	} else {
		if v, err := AppConfig.Int("httpport"); err == nil {
			HttpPort = v
		}

	}
	return nil
}
