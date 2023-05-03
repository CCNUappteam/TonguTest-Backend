package service

import (
	U "tongue/model/user"
	"tongue/pkg/errno"
)

// 更新用户信息

func UpdateInfo(email string, avatar string, name string, age string, phone string) error {
	if err := U.UpdateInfo(email, avatar, name, age, phone); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

// 修改密码
func UpdatePassword(email string, original_password string, new_password string) error {
	if err := U.UpdatePassword(email, original_password, new_password); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
