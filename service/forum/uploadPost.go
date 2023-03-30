package service

import (
	"tongue/model/forum"
	"tongue/pkg/errno"
)

func UploadPost(email, title, content string) (int, error) {
	post_id, err := forum.CreatePost(email, title, content)
	if err != nil {
		return post_id, errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return post_id, nil
}

func UploadPostImage(id int, url string) error {
	if err := forum.UploadPostImage(uint(id), url); err != nil {
		return errno.ServerErr(errno.ErrUploadFile, err.Error())
	}
	return nil
}
