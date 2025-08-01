/*
响应函数的封装
*/

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"-"`              // 响应码, 返回的json中忽略该字段
	Code   int    `json:"code,omitempty"` // 错误码
	Msg    string `json:"msg,omitempty"`  // 消息信息
	Data   any    `json:"data,omitempty"` // 响应数据
}

func (r ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(r, ResponseJson{})
}

func buildResponseFunc(status int) func(c *gin.Context, resp ResponseJson) {
	return func(c *gin.Context, resp ResponseJson) {
		if resp.IsEmpty() {
			c.AbortWithStatus(status)
		}
		if resp.Status == 0 {
			c.AbortWithStatusJSON(status, resp)
		} else {
			c.AbortWithStatusJSON(resp.Status, resp)
		}
	}
}

var Success = buildResponseFunc(http.StatusOK)
var ClientFail = buildResponseFunc(http.StatusBadRequest)
var ServerFail = buildResponseFunc(http.StatusInternalServerError)
