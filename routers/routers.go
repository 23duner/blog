package routers

import (
	"github.com/gin-gonic/gin"
	"hw_blog0/controller"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //设置成发布模式
	r := gin.New()
	v := r.Group("api/007")
	//注册登陆业务
	v.POST("/login", controller.LoginHandler)
	v1 := r.Group("user")
	v1.POST(" /blog/selectPage", controller.SelectPage)
	v1.POST("/user/add", controller.Add)
	v1.PUT("/user/update", controller.Update)
	v1.PUT("/updatePassword", controller.UpdatePassword)
	v1.DELETE("/user/delete", controller.Delete)
	return r
}
