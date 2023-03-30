package service

import (
	"tongue/model/forum"
	"tongue/pkg/errno"
)

func PostComment(email string, id uint, content string) error {
	if err := forum.CreateComment(email, id, content); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}

func DeleteComment(email, id string) error {
	if err := forum.DeleteComment(email, id); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}

func GetComments(id string) ([]*forum.Comment, error) {
	comments, err := forum.GetComments(id)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return comments, nil
}
