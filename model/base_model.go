package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"midnightapisvr/config"
)

const MYSQL_TIME_OFFSET = 28800

var midDBConfigStr string

func init() {
	midDBConfig := config.MidnightDBConfig
	midDBConfigStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		midDBConfig.User, midDBConfig.Password,
		midDBConfig.Host, midDBConfig.Port, midDBConfig.Database)
}

// GetMidnightDBConnection 获取 Midnight MySQL DB 的 连接
func GetMidnightDBConnection() (*sql.DB, error) {
	conn, err := sql.Open("mysql", midDBConfigStr)
	return conn, err
}
