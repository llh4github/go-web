package common

import "gorm.io/gorm"

// 全局变量
// 本项目使用依赖注入框架感觉上用处不太大，不如用全局变量来的简单方便
var (
	// MyDB 当前数据库连接实例
	MyDB *gorm.DB
	// MyConf 项目配置信息
	MyConf Conf
)

// -------------------------- public struct ------------------------------

// Conf 所有配置项
type Conf struct {
	// URLPrefix gin会判断是不是以 / 开头，不用考虑gin配置。
	// 为了使用方便请以 / 开头配置此项目
	URLPrefix string

	// App的名字，暂时没什么用。
	Name string
	// 启动端口号，默认8080
	Port string
	// Dbconfig 数据库配置
	Dbconfig DBConfig
	// Jwt jwt配置
	Jwt JwtConfig
}

// DBConfig 数据库连接配置
type DBConfig struct {
	Username, Password, Host, Dbname, Params string
	// Debug 是否开启gorm的debug模式
	Debug bool
}

// JwtConfig 生成jwt的配置
type JwtConfig struct {
	Secret string
	Iss    string
	// Exp 过期时间（分钟）。务必写入正整数
	Exp int
}
