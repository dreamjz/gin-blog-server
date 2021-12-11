package global

import (
	"gin-blog-server/models/config"

	"github.com/spf13/viper"
)

var (
	AppCfg   config.AppCfg
	AppViper *viper.Viper
)
