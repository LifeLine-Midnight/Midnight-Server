package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	MD5_SALT = "orianna#"
)

type PasswordUtils struct{}

func (*PasswordUtils) GetSaltPasswordHash(password string) string {
	passwordSalt := password + MD5_SALT
	h := md5.New()
	h.Write([]byte(passwordSalt))
	passwordHash := hex.EncodeToString(h.Sum(nil))

	return passwordHash
}
