package global

import (
	"go.uber.org/zap"
	"gogofly/conf"
	"gorm.io/gorm"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *conf.RedisClient
)
