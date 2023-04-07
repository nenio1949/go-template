package global

import (
	"go-template/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	TokenType         = "bearer"
	AppGuardName      = "system"
	WechatRedirectUrl = "https://www.baidu.com"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
