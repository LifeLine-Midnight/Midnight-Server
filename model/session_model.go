package model

// 用户会话数据访问层
type SessionModel struct{}

// CreateUserSession 创建用户会话
func (*SessionModel) CreateUserSession(uid int, token string) error {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	var ssql = `REPLACE INTO user_session(uid, stoken)
				VALUES(?, ?)`
	_, err = conn.Exec(ssql, uid, token)
	if err != nil {
		return err
	}

	return nil
}

// GetUidByToken 通过 Token 获取 Uid
func (*SessionModel) GetUidByToken(token string) (int, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	var ssql = `SELECT uid FROM user_session
				WHERE stoken=?`
	var uid int
	err = conn.QueryRow(ssql, token).Scan(&uid)
	if err != nil {
		return -1, err
	}

	return uid, nil
}

// DeleteUserSession 删除会话
func (*SessionModel) DeleteUserSession(token string) error {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	var ssql = `DELETE FROM user_session WHERE stoken=?`
	_, err = conn.Exec(ssql, token)
	if err != nil {
		return err
	}

	return nil
}
