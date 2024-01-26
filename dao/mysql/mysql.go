package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hw_blog0/config"
)

var db *sqlx.DB

// Init 初始化MySQL连接
func Init(cfg *config.Mysql) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenCons)
	db.SetMaxIdleConns(cfg.MaxIdleCons)
	return
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
