package forum

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"tongue/model"
)

type Like struct {
	gorm.Model
	PostId string
	Email  string
}

func CreateLike(email string, id string) error {
	var like Like
	if err := model.DB.Self.Model(&Like{}).Where("email = ? and post_id = ?", email, id).Find(&like).Error; err == nil {
		return nil
	}

	like = Like{
		PostId: id,
		Email:  email,
	}

	if err := model.DB.Self.Create(&like).Error; err != nil {
		return err
	}

	var post Post
	if err := model.DB.Self.Model(&Post{}).Where("id = ?", id).Find(&post).Error; err != nil {
		return err
	}
	post.LikeNum += 1
	if err := model.DB.Self.Model(&Post{}).Update(&post).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLike(email, id string) error {
	var like Like
	if err := model.DB.Self.Model(&Like{}).Where("post_id = ?", id).Find(&like).Error; err != nil {
		return err
	}
	fmt.Println(like.Email)
	if like.Email != email {
		return errors.New("permission denied")
	}
	if err := model.DB.Self.Model(&Like{}).Where("post_id = ?", id).Delete(&like).Error; err != nil {
		return err
	}
	return nil
}

func GetLikes(postid string) (int, error) {
	var count int
	if err := model.DB.Self.Model(&Like{}).Where("post_id = ?", postid).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}
