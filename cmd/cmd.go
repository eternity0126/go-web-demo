package cmd

import (
	"fmt"
	"gogofly/conf"
	"gogofly/global"
	"gogofly/router"
	"gogofly/utils"
)

func Start() {
	var initErr error
	// 读取settings.yml
	conf.InitConfig()

	// 初始化日志组件
	global.Logger = conf.InitLogger()

	// 初始化数据库连接
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// 初始化Redis连接
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	//_ = global.RedisClient.Set("username", "zs")
	//fmt.Println(global.RedisClient.Get("username"))

	// 初始化系统路由
	router.InitRouter()

	// 初始化过程中错误
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}

		panic(initErr.Error())
	}
}

func Clean() {
	fmt.Println("======Clean======")
}
