package router

import (
	"github.com/gin-gonic/gin"
	"gogofly/api"
	"net/http"
)

func InitUserRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.GET("", func(c *gin.Context) {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zs"},
					},
				})
			})

			rgAuthUser.GET("/:id", func(c *gin.Context) {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "zs",
				})
			})
		}
	})
}
