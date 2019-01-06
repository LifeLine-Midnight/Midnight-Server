package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"midnightapisvr/comm"
	"midnightapisvr/model"
	"midnightapisvr/utils"
)

// 用户会话服务
type SessionService struct{}

// UserLogIn 用户登录，创建会话项，返回 token
func (sessionSvc *SessionService) UserLogIn(username, password string) (*comm.SessionInfo, error) {
	var userModel = new(model.UserModel)

	userConfirmInfo, err := userModel.GetUserConfirmInfo(username)
	if err != nil {
		return nil, err
	}

	var pswdUtils = new(utils.PasswordUtils)
	pswdHash := pswdUtils.GetSaltPasswordHash(password)
	if pswdHash != userConfirmInfo.PasswordHash {
		return nil, errors.New("password error")
	}

	token := sessionSvc.generateToken(username)

	var sessionModel = new(model.SessionModel)
	err = sessionModel.CreateUserSession(userConfirmInfo.Uid, token)
	if err != nil {
		return nil, err
	}

	var sessionInfo = &comm.SessionInfo{
		Username: username,
		Token:    token,
	}

	return sessionInfo, nil
}

// UserLogOut 用户登出，删除会话项
func (*SessionService) UserLogOut(token string) error {
	var sessionModel = new(model.SessionModel)
	err := sessionModel.DeleteUserSession(token)
	if err != nil {
		return err
	}

	return nil
}

// GetUidByToken 通过 Token 获取 Uid
func (*SessionService) GetUidByToken(token string) (int, error) {
	var sessionModel = new(model.SessionModel)
	uid, err := sessionModel.GetUidByToken(token)
	if err != nil {
		return -1, err
	}

	return uid, nil
}

func (*SessionService) generateToken(username string) string {
	h := md5.New()
	h.Write([]byte(username))
	usernameHash := hex.EncodeToString(h.Sum(nil))
	token := fmt.Sprintf("%s%d", usernameHash, time.Now().UnixNano())

	return token
}
