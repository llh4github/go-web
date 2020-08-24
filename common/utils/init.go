package utils

import "gitee.com/llh-gitee/go-web/config"

// jwtConf 生成jwt的配置
var jwtConf config.JwtConfig

func init() {
	jwtConf = config.GetJwtConfig()
}
