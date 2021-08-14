package main

import (
	"api-go-vue-admin/controller/admin"
	"api-go-vue-admin/controller/article"
	"api-go-vue-admin/controller/category"
	"api-go-vue-admin/controller/search"
	"api-go-vue-admin/middleware"
	"api-go-vue-admin/router"
)

func main() {

	// 加载多个APP的路由配置  多个样例：router.Include(admin.Routers1,admin.Router2,admin.Router3)
	router.Include(admin.Routers, category.Routers, article.Routers, search.Routers)
	// 初始化路由
	r := router.Init()
	//使用跨域中间件
	r.Use(middleware.Cors())
	//绑定程序运行端口
	r.Run(":8082")

}
