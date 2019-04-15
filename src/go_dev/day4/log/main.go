package main

import "encoding/json"
import "fmt"
import "github.com/astaxie/beego/logs"

func main(){
	config := make(map[string]interface{})
	config["fileanme"] = "./logs/logcollect.log"
	config["level"] = logs.LevelDebug
	configStr,err := json.Marshal(config)
	if err != nil{
		fmt.Println("marshal failed err:",err)
		return
	}
	logs.SetLogger(logs.AdapterFile,string(configStr))
	logs.Debug("This is a test my name is %s","stu01")
	logs.Trace("This is a trace my name is %s","stu02")
	logs.Warn("This is a warn my name is %s","stu03")
}