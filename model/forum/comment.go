package forum

import (
	"errors"
	"github.com/jinzhu/gorm"
	"tongue/model"
)

type Comment struct {
	gorm.Model
	PublisherEmail string `json:"publisher_email" gorm:"column:publisher_id"`
	PostId         uint   `json:"post_id" gorm:"column:post_id"`
	Content        string `json:"content" gorm:"column:content"`
}

func CreateComment(email string, id uint, content string) error {
	var comment = Comment{
		PublisherEmail: email,
		PostId:         id,
		Content:        content,
	}
	if err := model.DB.Self.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteComment(email, id string) error {
	var comment Comment
	if err := model.DB.Self.Model(&Comment{}).Where("id = ?", id).Find(&comment).Error; err != nil {
		return err
	}
	if comment.PublisherEmail != email {
		return errors.New("permission denied")
	}
	if err := model.DB.Self.Where("id = ?", id).Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

func GetComments(id string) ([]*Comment, error) {
	comments := make([]*Comment, 0)
	if err := model.DB.Self.Model(&Comment{}).Where("post_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
