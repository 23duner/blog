package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hw_blog0/dao/mysql"
	"hw_blog0/global"
	"hw_blog0/logic"
	"hw_blog0/models"
)

func LoginHandler(c *gin.Context) {
	//获取请求参数及参数校验
	var u *models.LoginForm
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误，直接返回响应
		global.Log.Error("Login with invalid param")
		ResponseError(c, CodeInvalidParams) // 请求参数错误
		return
	}
	//处理业务逻辑-登录
	user, err := logic.Login(u)
	if err != nil {
		global.Log.Error("logic.Login failed")
		if err.Error() == mysql.ErrorUserNotExit {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.ID), //js识别的最大值：id值大于1<<53-1  int64: i<<63-1
		"user_name":     user.UserName,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}
func SelectPage(c *gin.Context) {
	var p *models.Page
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误，直接返回响应
		global.Log.Error("selectPage with invalid param")
		ResponseError(c, CodeInvalidParams) // 请求参数错误
		return
	}
	blogs, err := logic.SelectPage(p)
	if err != nil {
		global.Log.Error("logic.selectPage failed", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, blogs)
}
func Add(c *gin.Context) {
	//获取请求参数及参数校验
	var u *models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误，直接返回响应
		global.Log.Error("Add with invalid param")
		ResponseError(c, CodeInvalidParams) // 请求参数错误
		return
	}
	//处理业务逻辑-新增
	if err := logic.Add(u); err != nil {
		global.Log.Error("logic.adduser failed", err)
		if err.Error() == mysql.ErrorUserExit {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
func Update(c *gin.Context) {
	//获取请求参数及参数校验
	var u *models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误，直接返回响应
		global.Log.Error("update with invalid param")
		ResponseError(c, CodeInvalidParams) // 请求参数错误
		return
	}
	//处理业务逻辑-更新用户
	if err := logic.Update(u); err != nil {
		global.Log.Error("logic.update failed", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
func Delete(c *gin.Context) {
	//获取请求参数及参数校验
	var id int
	if err := c.ShouldBindJSON(&id); err != nil {
		//请求参数有误，直接返回响应
		global.Log.Error("delete with invalid param")
		ResponseError(c, CodeInvalidParams) // 请求参数错误
		return
	}
	//处理业务逻辑-删除用户
	if err := logic.Delete(id); err != nil {
		global.Log.Error("logic.delete failed", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
func UpdatePassword(c *gin.Context) {
	//获取请求参数及参数校验
	var u *models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		//请求参数有误，直接返回响应
		global.Log.Error("updatePassword with invalid param")
		ResponseError(c, CodeInvalidParams) // 请求参数错误
		return
	}
	//处理业务逻辑-更新密码
	if err := logic.UpdatePassword(u); err != nil {
		global.Log.Error("logic.updatePassword failed", err)
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
