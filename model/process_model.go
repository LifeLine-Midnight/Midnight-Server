package model

import (
	"time"

	"midnightapisvr/comm"
	"midnightapisvr/utils"
)

type ProcessModel struct{}

// CreateUserProcess 初始化一份用户进度 item
func (*ProcessModel) CreateUserProcess(uid int) error {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	var ssql = `INSERT INTO user_process(uid) VALUES(?)`
	_, err = conn.Exec(ssql, uid)
	if err != nil {
		return err
	}

	return nil
}

func (*ProcessModel) GetUserProcess(uid int) (*comm.UserProcessDB, error) {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var ssql = `SELECT uid, cur_sid, cur_process_time FROM user_process
				WHERE uid=?`
	var userProcessInfo = new(comm.UserProcessDB)
	var exeTime time.Time

	err = conn.QueryRow(ssql, uid).Scan(&userProcessInfo.Uid,
		&userProcessInfo.CurSid, &exeTime)
	if err != nil {
		return nil, err
	}

	userProcessInfo.CurProcessTime = exeTime.Unix() - MYSQL_TIME_OFFSET
	return userProcessInfo, nil
}

func (*ProcessModel) UpdateUserProcess(uid int, curSid int, curProcessTime int64) error {
	conn, err := GetMidnightDBConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	var ssql = `UPDATE user_process SET cur_sid=?, cur_process_time=?
				WHERE uid=?`

	var timeUtils = new(utils.TimeUtils)
	curProcessTimeStr := timeUtils.TimestampToStr(curProcessTime)
	_, err = conn.Exec(ssql, curSid, curProcessTimeStr, uid)
	if err != nil {
		return err
	}

	return nil
}
