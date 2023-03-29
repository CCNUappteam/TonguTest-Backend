package service

import (
	"fmt"
	user "tongue/model/user"
	"tongue/pkg/errno"
)

// getInfo ...获取个人信息
func GetInfo(email string) (*user.UserModel, error) {
	userModel := &user.UserModel{}
	err := userModel.GetInfo(email)
	fmt.Println("user----", userModel)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return userModel, nil
}
