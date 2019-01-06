package config

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     uint32
	Database string
}

var MidnightDBConfig *MySQLConfig

func init() {
	MidnightDBConfig = &MySQLConfig{
		User:     "user",
		Password: "user999",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "midnight",
	}
}
