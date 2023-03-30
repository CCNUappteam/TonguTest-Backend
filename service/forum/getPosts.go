package service

import (
	"tongue/model/forum"
	"tongue/pkg/errno"
)

func GetPosts(email string, offset, limit int) ([]*forum.Post, int, error) {
	post, count, err := forum.GetPost(email, offset, limit)
	if err != nil {
		return nil, 0, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return post, count, nil
}
