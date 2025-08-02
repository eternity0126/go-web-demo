package service

import (
	"errors"
	"gogofly/dao"
	"gogofly/model"
	"gogofly/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		return &UserService{
			Dao: dao.NewUserDao(),
		}
	}

	return userService
}

func (u *UserService) Login(iUserDTO dto.UserLoginDTO) (model.User, error) {
	var errResult error

	iUser := u.Dao.GetUserByNameAndPassword(iUserDTO.Name, iUserDTO.Password)
	if iUser.ID == 0 {
		errResult = errors.New("Invalid Username or Password")
	}

	return iUser, errResult
}

func (u *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if u.Dao.CheckUsernameExist(iUserAddDTO.Name) {
		return errors.New("Username exists")
	}
	return u.Dao.AddUser(iUserAddDTO)
}

func (u *UserService) GetUserById(iCommonIdDTO *dto.CommonIdDTO) (model.User, error) {
	return u.Dao.GetUserById(iCommonIdDTO.ID)
}
