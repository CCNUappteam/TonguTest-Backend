package service

import (
	"tongue/model/user"
	"tongue/pkg/errno"
)

func UpdateAvatar(email, url string) error {
	if err := user.UpdateInfo(email, url, ""); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
