package service

import (
	"midnightapisvr/comm"
	"midnightapisvr/model"
	"midnightapisvr/utils"
)

type UserService struct{}

// UserRegister 用户注册
func (userSvc *UserService) UserRegister(username, nickname,
	password, avatarURI string) (*comm.UserInfoRes, error) {
	// 加盐哈希密码
	var pswdUtils = new(utils.PasswordUtils)
	passwordHash := pswdUtils.GetSaltPasswordHash(password)

	var userModel = new(model.UserModel)
	uid, err := userModel.CreateUser(username, nickname, passwordHash, avatarURI)
	if err != nil {
		return nil, err
	}

	// 初始化进度
	var processModel = new(model.ProcessModel)
	err = processModel.CreateUserProcess(uid)
	if err != nil {
		return nil, err
	}

	var userInfo = &comm.UserInfoRes{
		Username:  username,
		Nickname:  nickname,
		AvatarURI: avatarURI,
	}

	return userInfo, nil
}

// GetUserInfo 通过 uid 获取用户详细信息
func (*UserService) GetUserInfo(uid int) (*comm.UserInfoRes, error) {
	var userModel = new(model.UserModel)
	userInfoDB, err := userModel.GetUserInfo(uid)
	if err != nil {
		return nil, err
	}

	var userInfo = &comm.UserInfoRes{
		Username:  userInfoDB.Username,
		Nickname:  userInfoDB.Nickname,
		AvatarURI: userInfoDB.AvatarURI,
	}

	return userInfo, nil
}
