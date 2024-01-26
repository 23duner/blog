package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"hw_blog0/models"
)

const secret = "007.vip"

// encryptPassword 对密码进行加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}
func Login(user *models.User) (err error) {
	originPassword := user.Password // 记录一下原始密码(用户登录的密码)
	sqlStr := "select id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	// 查询数据库出错
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	// 用户不存在
	if err == sql.ErrNoRows {
		return errors.New(ErrorUserNotExit)
	}
	// 生成加密密码与查询到的密码比较
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return errors.New(ErrorPasswordWrong)
	}
	return nil
}

// 新增用户
func AddUser(p models.User) (error error) {
	//执行sql语句入库
	sqlstr := `insert into user(username,password,name,phone,email,avatar,role,sex,info,birth) values(?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlstr, p.UserName, p.Password, p.Name, p.Phone, p.Email, p.Avatar, p.Role, p.Sex, p.Info, p.Info, p.Birth)
	return err
}

// 查询用户
func SelectByUsername(name string) (error error) {
	// 编写SQL查询语句，检查users表中是否存在相同的用户名
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", name).Scan(&name)
	if err != nil {
		// 如果查询过程中发生错误，返回错误
		return err
	}
	// 如果exists为true，说明用户名已存在
	return nil
}
