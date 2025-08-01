package conf

import (
	"github.com/spf13/viper"
	"gogofly/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if viper.GetString("env") != "dev" {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.max_idle_conn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.max_open_conn"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{})
	return db, nil
}
