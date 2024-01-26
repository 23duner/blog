package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hw_blog0/core"
	"hw_blog0/global"
	"hw_blog0/routers"
)

func main() {
	//读取配置文件
	core.Init()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		global.Log.Error(err)
	}
	defer db.Close()
	//注册路由
	r := routers.Init()
	err = r.Run(fmt.Sprintf(":%d", global.Config.System.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
