package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gogofly/docs"
	"gogofly/global"
	"gogofly/middleware"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

/*
路由初始化
*/

type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRoute
)

func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitRouter() {
	// 创建监听中断信号, 应用退出信号的上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	r := gin.Default()
	r.Use(middleware.Cors())
	// 设置两个路由组
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	// 初始化基础平台的路由
	initBasePlatformRoutes()

	// 注册自定义验证器
	registerValidator()

	// 注册路由
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	// 继承swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		global.Logger.Infof("Start listening: %s", stPort)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Errorf("Start server error: %s\n", err.Error())
			return
		}
	}()

	// 等待通道
	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Errorf("Stop server error: %s\n", err.Error())
		return
	}

	global.Logger.Info("Stop server success")
}

func initBasePlatformRoutes() {
	InitUserRoutes()
	InitHostRoutes()
}

// 注册自定义验证器
func registerValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && 0 == strings.Index(value, "a") {
					return true
				}
			}

			return false
		})
	}
}
