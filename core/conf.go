package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"hw_blog0/config"
	"hw_blog0/global"
	"io/ioutil"
	"log"
)

// Init 读取yaml文件的配置
func Init() {
	const ConfigFile = "settings.yaml" //yaml文件路径
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile) //读取yaml文件
	if err != nil {
		panic(fmt.Errorf("get yamlConf error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c) //若解析失败
	if err != nil {
		log.Fatalf("config  Init Unmarshal:%v", err) //此时yaml文件还未完全初始化，所以只能用原生log不能用自己的
	}
	log.Println("config yamlFile load Init success")
	global.Config = c

}
