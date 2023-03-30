package service

import (
	"tongue/model/forum"
	"tongue/pkg/errno"
)

func GetPostImage(id string) ([]*forum.PostImage, error) {
	images, err := forum.GetPostImage(id)
	if err != nil {
		return nil, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return images, nil
}
