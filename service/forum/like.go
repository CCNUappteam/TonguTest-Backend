package service

import (
	"tongue/model/forum"
	"tongue/pkg/errno"
)

func Like(email, id string) error {
	if err := forum.CreateLike(email, id); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func DeleteLike(email, id string) error {
	if err := forum.DeleteLike(email, id); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func GetLikes(id string) (int, error) {
	count, err := forum.GetLikes(id)
	if err != nil {
		return -1, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return count, nil
}
