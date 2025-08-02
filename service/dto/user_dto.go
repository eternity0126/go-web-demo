package dto

import "gogofly/model"

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"Username can not be empty"`
	Password string `json:"password" binding:"required" message:"Password can not be empty"`
}

// 添加用户的DTO
type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"Username can not be empty"`
	Avatar   string
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"Password can not be empty"`
}

// 将DTO接收到的数据传递到数据库实体上
func (u *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = u.Name
	iUser.Avatar = u.Avatar
	iUser.Phone = u.Phone
	iUser.Email = u.Email
	iUser.Password = u.Password
}
