package global

import (
	"github.com/sirupsen/logrus"
	"hw_blog0/config"
)

var (
	Config *config.Config //需要一个全局变量存放在global目录下，用于保存配置文件
	Log    *logrus.Logger
)
