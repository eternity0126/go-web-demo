package api

import (
	"github.com/gin-gonic/gin"
	"gogofly/service"
	"gogofly/service/dto"
)

type HostApi struct {
	BaseApi
	Service *service.HostService
}

func NewHostApi() HostApi {
	return HostApi{
		Service: service.NewHostService(),
	}
}

// TODO: 为什么不能用指针？
func (h HostApi) Shutdown(c *gin.Context) {
	var iShutdownHostDTO dto.ShutdownHostDTO
	if err := h.BuildRequest(BuildRequestOption{Ctx: c, DTO: iShutdownHostDTO}); err != nil {
		return
	}

	err := h.Service.Shutdown(iShutdownHostDTO)

	if err != nil {
		h.ClientFail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
	}

	h.Success(ResponseJson{
		Msg: "Shutdown success",
	})
}
