package tool

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var Logs *logs.BeeLogger

func init() {
	Logs = setlog()
}

type logConf struct {
	Filename string `json:"filename"`
	Maxdays  int64  `json:"maxdays"`
	Maxsize  int64  `json:"maxsize"`
	Rotate   bool   `json:"rotate"`
	MaxLines int64  `json:"maxlines"`
}

func setlog() *logs.BeeLogger {
	log := logs.NewLogger(10000)
	filename := beego.AppConfig.String("logfilename")
	maxdays, _ := beego.AppConfig.Int64("maxdays")
	maxsize, _ := beego.AppConfig.Int64("maxsize")
	rotate, _ := beego.AppConfig.Bool("rotate")
	maxlines, _ := beego.AppConfig.Int64("maxlines")
	logConf := logConf{
		Filename: filename,
		Maxdays:  maxdays,
		Maxsize:  maxsize,
		Rotate:   rotate,
		MaxLines: maxlines,
	}
	b, _ := json.Marshal(logConf)
	log.SetLogger(logs.AdapterFile, string(b))
	log.EnableFuncCallDepth(true)
	log.Info(string(b))
	//log.Flush()
	return log
}
