package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"hw_blog0/models"
)

const secret = "007.vip"

// encryptPassword 对密码进行加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

// Login 检测登录信息是否匹配
func Login(user *models.User) (err error) {
	_, err = db.Exec("select * from users where username = ? and password = ?", user.UserName, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// AddUser 新增用户
func AddUser(p models.User) (error error) {
	//执行sql语句入库
	sqlstr := `insert into users(username,password,name,phone,email,avatar,role,sex,info,birth) values(?,?,?,?,?,?,?,?,?,?)`
	_, err := db.Exec(sqlstr, p.UserName, p.Password, p.Name, p.Phone, p.Email, p.Avatar, p.Role, p.Sex, p.Info, p.Info, p.Birth)
	return err
}

// SelectByUsername 查询用户
func SelectByUsername(name string) (err error) {
	rows, err := db.Query("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

// SelectPage 分页查询博客
func SelectPage(p *models.Page) ([]models.Blog, error) {
	rows, err := db.Query(`SELECT * FROM blog b JOIN (SELECT id, name FROM users) u ON b.user_id = u.id WHERE b.content LIKE ? LIMIT ?, ?`, p.Text, p.PageNum, p.PageSize)
	if err != nil {
		return nil, err
	}

	var results []models.Blog
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Descr, &blog.Cover, &blog.Tags, &blog.UserId, &blog.Name, &blog.Date, &blog.ReadCount, &blog.CategoryId)
		if err != nil {
			rows.Close()
			return nil, err
		}
		results = append(results, blog)
	}
	err = rows.Err()
	if err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()
	return results, nil
}

func Update(user *models.User) (err error) {
	// 执行更新操作
	sqlstr := "UPDATE users SET username = ?,  name = ?,  avatar = ?,  role = ?,  sex = ?,  phone = ?,  email = ?,  info = ?,  birth = ?  WHERE  username = ?;"
	_, err = db.Exec(sqlstr, user.UserName, user.Name, user.Avatar, user.Role, user.Sex, user.Phone, user.Email, user.Info, user.Birth, user.UserName)
	if err != nil {
		return err
	}
	return nil
}
func UpdatePassword(user *models.User) (err error) {
	// 执行更新操作
	_, err = db.Exec("update users set password = ? where username = ?", user.Password, user.UserName)
	if err != nil {
		return err
	}
	return nil
}

func Delete(id int) (err error) {
	_, err = db.Exec("Delete * from users where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// SelectUser 根据用户查询博客
func SelectUser(p *models.Page) ([]models.Blog, error) {
	rows, err := db.Query(`SELECT * FROM blog WHERE username = ?`, p.Username)
	if err != nil {
		return nil, err
	}

	var results []models.Blog
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Descr, &blog.Cover, &blog.Tags, &blog.UserId, &blog.Name, &blog.Date, &blog.ReadCount, &blog.CategoryId)
		if err != nil {
			rows.Close()
			return nil, err
		}
		results = append(results, blog)
	}
	err = rows.Err()
	if err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()
	return results, nil
}

// SelectLike 根据点赞查询博客
func SelectLike(p *models.Page) ([]models.Blog, error) {
	rows, err := db.Query(`SELECT * FROM blog WHERE id IN (SELECT fid FROM likes WHERE user_name = ？)`, p.Username)
	if err != nil {
		return nil, err
	}

	var results []models.Blog
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Descr, &blog.Cover, &blog.Tags, &blog.UserId, &blog.Name, &blog.Date, &blog.ReadCount, &blog.CategoryId)
		if err != nil {
			rows.Close()
			return nil, err
		}
		results = append(results, blog)
	}
	err = rows.Err()
	if err != nil {
		rows.Close()
		return nil, err
	}
	rows.Close()
	return results, nil
}
