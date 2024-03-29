package models

// User 定义请求参数结构体
type User struct {
	ID           uint64 `json:"id" db:"id"`             // 指定json序列化/反序列化时使用小写user_id
	UserName     string `json:"username" db:"username"` //用户名
	Password     string `json:"password" db:"password"`
	Name         string `json:"name" db:"name"`     //姓名
	Avatar       string `json:"avatar" db:"avatar"` //头像
	Role         string `json:"role" db:"role"`     //角色标识
	Sex          string `json:"sex" db:"sex"`       //性别
	Phone        string `json:"phone" db:"phone"`   //电话
	Email        string `json:"email" db:"email"`   //邮箱
	Info         string `json:"info" db:"info"`     //简介
	Birth        string `json:"birth" db:"birth"`   //生日
	AccessToken  string
	RefreshToken string
}

// LoginForm  登录请求参数
type LoginForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Page struct {
	PageNum  int    `json:"pageNum" binding:"required"`
	PageSize int    `json:"pageSize" binding:"required"`
	Text     string `json:"title"`
}
type Add struct {
	UserName string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phine" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Info     string `json:"info" binding:"required"`
	Birth    string `json:"birth" binding:"required"`
}
type Blog struct {
	ID         uint64 `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Content    string `json:"content" db:"content"`
	Descr      string `json:"descr" db:"descr"`
	Cover      string `json:"cover" db:"cover"`
	Tags       string `json:"tags" db:"tags"`
	UserId     string `json:"user_id" db:"user_id"`
	Name       string `json:"name"`
	Date       string `json:"date" db:"date"`
	ReadCount  string `json:"read_count" db:"read_count"`
	CategoryId string `json:"category_id" db:"category_id"`
}
