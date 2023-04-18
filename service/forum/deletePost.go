package service

import (
	"tongue/model/forum"
	"tongue/pkg/errno"
)

func DeletePost(email, id string) error {
	if err := forum.DeletePost(email, id); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
