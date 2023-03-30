package forum

import (
	"errors"
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
	if err := model.DB.Self.Model(&Like{}).Where("email = ? and post_id = ?", email, id).Find(&like).Error;
		err == nil {
		return nil
	}

	like = Like{
		PostId: id,
		Email:  email,
	}

	if err := model.DB.Self.Create(&like).Error; err != nil {
		return err

	}
	return nil
}

func DeleteLike(email, id string) error {
	var like Like
	if err := model.DB.Self.Model(&Like{}).Where("id = ?", id).Find(&like).Error; err != nil {
		return err
	}
	if like.Email != email {
		return errors.New("permission denied")
	}
	if err := model.DB.Self.Model(&Like{}).Where("id = ?", id).Delete(&like).Error; err != nil {
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
