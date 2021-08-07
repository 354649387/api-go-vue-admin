package main

import (
	"api-go-vue-admin/middleware"
	"api-go-vue-admin/router"
)

func main(){


	//设置路由
	r := router.SetRouter()
	//使用跨域中间件
	r.Use(middleware.Cors())
	//绑定程序运行端口
	r.Run(":8082")

}