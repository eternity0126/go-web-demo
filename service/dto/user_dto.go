package dto

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"用户名填写错误" required_err:"Username can not be empty"`
	Password string `json:"password" binding:"required" message:"Password can not be empty"`
}
