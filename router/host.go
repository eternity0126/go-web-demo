package router

import (
	"github.com/gin-gonic/gin"
	"gogofly/api"
)

func InitHostRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		hostApi := api.NewHostApi()

		rgAuthHost := rgAuth.Group("host")
		{
			rgAuthHost.GET("/shutdown", hostApi.Shutdown)
		}
	})
}
