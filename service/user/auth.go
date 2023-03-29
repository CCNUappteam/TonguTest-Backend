package service

import (
	"tongue/model/user"
	"tongue/pkg/errno"
)

// Authorize ... 授权
func Authorize(email string, role int) error {
	if err := user.Authorize(email, role); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// GetAuthority ...获取管理员信息
func GetAuthority(role int) ([]user.UserModel, error) {
	admin, err := user.GetAdimin(role)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return admin, nil
}
