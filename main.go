package main

import (
	"GoWebApp/controllers"
	"GoWebApp/dao/mysql"
	"GoWebApp/logger"
	"GoWebApp/routes"
	"GoWebApp/settings"
	"fmt"
)

func main() {

	//1:初始化配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("初始化配置文件失败:%v\n", err)
		return
	}

	//2:初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("初始化日志失败:%v\n", err)
		return
	}
	//3:初始化mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("初始化数据库失败:%v\n", err)
		return
	}

	//4:初始化redis
	//if err := redis.Init(); err != nil {
	//	fmt.Printf("初始化redis失败：%v\n", err)
	//	return
	//}

	//5:注册路由
	r := routes.SetUp()

	//6:优雅关机
	controllers.ShutDown(r)
}
