package main

import (
	"webapp/database"
	"webapp/router"
)
func main() {
	// 初始化数据表
	database.InitMysql()
	// 创建一个路由实例
	r := router.InitRouter()
    
	r.Run(":8888")
}
