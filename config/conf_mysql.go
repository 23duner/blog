package config

type Mysql struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	DB          string `yaml:"db"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	LogLevel    string `yaml:"log_level"`     //日志等级，debug就是输出全部sql，dev
	MaxOpenCons int    `yaml:"max_open_cons"` //支持同时连接数据库的端口数
	MaxIdleCons int    `yaml:"maxIdleCons"`   //数据库支持并发数
}
