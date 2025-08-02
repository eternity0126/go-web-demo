package api

import (
	"github.com/gin-gonic/gin"
	"gogofly/service"
	"gogofly/service/dto"
	"gogofly/utils"
)

const (
	ERR_CODE_ADD_USER = 10011 + iota
	ERR_CODE_GET_USER_BY_ID
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO

	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, err := u.Service.Login(iUserLoginDTO)
	if err != nil {
		u.ClientFail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	// 成功查找到对应用户
	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	u.Success(ResponseJson{
		Data: gin.H{
			"token": token,
			"iUser": iUser,
		},
	})
}

func (u UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	err := u.Service.AddUser(&iUserAddDTO)
	if err != nil {
		u.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	u.Success(ResponseJson{
		Data: iUserAddDTO,
	})

}

func (u UserApi) GetUserById(c *gin.Context) {
	var iCommonIdDTO dto.CommonIdDTO
	if err := u.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIdDTO, BindParamsFromUri: true}).GetError(); err != nil {
		return
	}

	iUser, err := u.Service.GetUserById(&iCommonIdDTO)
	if err != nil {
		u.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	u.Success(ResponseJson{
		Data: iUser,
	})
}
