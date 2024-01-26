package logic

import (
	"fmt"
	"hw_blog0/dao/mysql"
	"hw_blog0/models"
	"hw_blog0/pkg/jwt"
)

func Login(p *models.LoginForm) (user *models.User, error error) {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	//return jwt.GenToken(user.UserID,user.UserName)
	accessToken, refreshToken, err := jwt.GenToken(user.ID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}

func Add(u *models.User) (error error) {
	//业务方法
	//1.判断用户账号是否重复
	err := mysql.SelectByUsername(u.UserName)
	if err != nil {
		// 如果存在重复的用户名，返回相应的错误信息
		return fmt.Errorf("用户名已存在: %w", err)
	}

	//2.判断用户密码是不是空
	if u.Password == "" {
		return fmt.Errorf("密码不能为空")
	}
	//3.判断用户名称是不是空
	if u.UserName == "" {
		return fmt.Errorf("用户名称不能为空")
	}
	//4.默认用户角色
	//构造一个User实例
	p := models.User{
		UserName: u.UserName,
		Name:     u.Name,
		Phone:    u.Phone,
		Email:    u.Email,
		Avatar:   u.Avatar,
		Role:     u.Role,
		Sex:      u.Sex,
		Info:     u.Info,
		Birth:    u.Birth,
	}
	//保存进数据库
	return mysql.AddUser(p)
}
