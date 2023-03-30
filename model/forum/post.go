package forum

import (
	"errors"
	"github.com/jinzhu/gorm"
	"tongue/model"
)

type Post struct {
	gorm.Model
	PublisherEmail string `json:"publisher_email" gorm:"column:publisher_email" binding:"required"`
	Title          string `json:"title" gorm:"column:title" binding:"required"`
	Content        string `json:"content" gorm:"column:content" binding:"required"`
	LikeNum        int    `json:"like_num" gorm:"column:like_num"`
}

type PostImage struct {
	gorm.Model
	PostID uint   `json:"post_id" gorm:"column:post_id"`
	Url    string `json:"url" gorm:"column:url"`
}

func CreatePost(email, title, content string) (int, error) {
	var post = Post{
		PublisherEmail: email,
		Title:          title,
		Content:        content,
	}
	if err := model.DB.Self.Create(&post).Error; err != nil {
		return -1, err
	}
	return int(post.ID), nil
}

func GetPost(email string, offset int, limit int) ([]*Post, int, error) {
	item := make([]*Post, 0)
	var count int

	if email != "" {
		d := model.DB.Self.Model(&Post{}).
			Offset(offset).Limit(limit).Where("publisher_email = ?", email).
			Order("created_at desc").Find(&item)
		if d.Error != nil {
			return nil, 0, d.Error
		}
		model.DB.Self.Model(&Post{}).Where("publisher_email = ?", email).
			Count(&count)
	} else {
		d := model.DB.Self.Model(&Post{}).
			Offset(offset).Limit(limit).
			Order("created_at desc").Find(&item)
		if d.Error != nil {
			return nil, 0, d.Error
		}
		model.DB.Self.Model(&Post{}).Count(&count)
	}
	return item, count, nil
}

func UploadPostImage(postId uint, url string) error {
	var post Post
	if err := model.DB.Self.Model(&Post{}).Where("id = ?", postId).Find(&post).Error; err != nil {
		return errors.New("没有该推文")
	}
	var postImage PostImage
	postImage.PostID = postId
	postImage.Url = url
	if err := model.DB.Self.Create(&postImage).Error; err != nil {
		return err
	}
	return nil
}

func GetPostImage(id string) ([]*PostImage, error) {
	images := make([]*PostImage, 0)
	if err := model.DB.Self.Model(&PostImage{}).Where("post_id = ?", id).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}
