package main

/*********************
**材料证明信息读取及展示
**提供材料证明的word下载功能
*********************/

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	controllers "github.com/dcmsy/cer_site/controllers"

	"github.com/donnie4w/go-logger/logger"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

//beego init
func init() {
	var config_file string
	flag.StringVar(&config_file, "conf", "", "the path of the config file")
	flag.Parse()
	if config_file != "" {
		beego.AppConfigPath, _ = filepath.Abs(config_file)
		beego.ParseConfig()
	} else {
		if config_file = os.Getenv("BEEGO_APP_CONFIG_FILE"); config_file != "" {
			beego.AppConfigPath, _ = filepath.Abs(config_file)
			beego.ParseConfig()
		}
	}
}

//main 主入口函数
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//指定是否控制台打印，默认为true
	logger.SetConsole(true)
	//指定日志级别  ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF 级别由低到高
	logger.SetLevel(logger.ERROR)
	logger.SetRollingFile("./log", "site.log", 100, 50, logger.MB)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	beego.Router("/", &controllers.IndexInfo{}, "*:Index")
	//启动参数管理UI 服务
	beego.Run()
	for {
		time.Sleep(time.Hour)
	}
}
