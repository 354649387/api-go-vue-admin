package router

import (
	adminController "api-go-vue-admin/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

//管理员列表
func adminList(c *gin.Context) {

	//接收结果
	admins := adminController.List()

	c.JSON(http.StatusOK,gin.H{
		"list":admins,
	})

}

//管理员新增
func adminAdd(c *gin.Context){

	res := adminController.Add()

	if res {
		c.JSON(http.StatusOK,"success")
	}

}

//文章列表
func articleList(c *gin.Context){

	c.String(200,"articleList")

}

//设置路由配置信息
func SetRouter() *gin.Engine {

	//创建一个路由
	r := gin.Default()

	//admin路由组
	admin := r.Group("/admin")
	{
		//管理员列表页
		admin.GET("/list", adminList)
		//管理员新增页面
		admin.GET("/add", adminAdd)

	}
	//article路由组
	article := r.Group("/article")
	{
		//文章列表页
		article.GET("/list",articleList)
	}

	return r
}
