package dao

import (
	"gogofly/model"
	"gogofly/service/dto"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}

	return userDao
}

// 查询用户信息
func (u *UserDao) GetUserByNameAndPassword(stUsername string, stPassword string) model.User {
	var iUser model.User
	u.Orm.Model(&iUser).Where("name=? and password=?", stUsername, stPassword).Find(&iUser)
	return iUser
}

func (u *UserDao) CheckUsernameExist(stUsername string) bool {
	var nTotal int64
	u.Orm.Model(&model.User{}).Where("name=?", stUsername).
		Count(&nTotal)

	return nTotal > 0
}

func (u *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)

	err := u.Orm.Save(&iUser).Error
	// 使用DTO实现客户端回显
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}

	return err
}

func (u *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := u.Orm.First(&iUser, id).Error
	return iUser, err
}
