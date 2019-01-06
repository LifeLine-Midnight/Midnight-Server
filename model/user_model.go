package model

import (
	"errors"

	"midnightapisvr/comm"
)

type UserModel struct{}

// CreateUser 创建一个用户 (玩家)
func (*UserModel) CreateUser(username, nickname, passwordHash, avatarURI string) (int, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var ssql = `INSERT INTO user_info (username, nickname, avatar_uri, passwordhash)
				VALUES(?, ?, ?, ?)`
	res, err := conn.Exec(ssql, username, nickname, avatarURI, passwordHash)
	if err != nil {
		return 0, err
	}
	uid, _ := res.LastInsertId()
	if uid <= 0 {
		return 0, errors.New("uid init error")
	}

	return int(uid), nil
}

// GetUserInfo 获取用户详细信息
func (*UserModel) GetUserInfo(uid int) (*comm.UserInfoDB, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT uid, username, nickname, avatar_uri
				FROM user_info WHERE uid=?`

	var userInfo = new(comm.UserInfoDB)
	err = conn.QueryRow(ssql, uid).Scan(&userInfo.Uid,
		&userInfo.Username, &userInfo.Nickname, &userInfo.AvatarURI)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

// GetUserPasswordHash 获取加盐哈希后的用户密码和uid
func (*UserModel) GetUserConfirmInfo(username string) (*comm.UserConfirmInfo, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT uid, passwordhash
				FROM user_info WHERE username=?`
	var userConfirmInfo = new(comm.UserConfirmInfo)
	err = conn.QueryRow(ssql, username).Scan(&userConfirmInfo.Uid, &userConfirmInfo.PasswordHash)
	if err != nil {
		return nil, err
	}

	return userConfirmInfo, nil
}
