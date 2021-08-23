package admin

import "github.com/gin-gonic/gin"

/**
管理员相关理由设置
*/

func Routers(e *gin.Engine) {

	admin := e.Group("/admin")
	{
		//管理员列表页
		admin.GET("/list", adminList)
		//管理员新增页面
		admin.GET("/add", adminAdd)
		//管理员搜索页面
		admin.GET("/search", adminSearch)
		//删除管理员
		admin.GET("/delete", adminDelete)
		admin.GET("/update", update)
		admin.GET("/getAdminById", getAdminById)
	}

}
