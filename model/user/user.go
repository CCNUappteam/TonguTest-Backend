package user

import (
	Md5 "crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
	"tongue/model"
)

type UserModel struct {
	gorm.Model
	Name         string `json:"name" gorm:"column:name;not null" binding:"required"`
	Email        string `json:"email" gorm:"column:email;default:null;unique"`
	Avatar       string `json:"avatar" gorm:"column:avatar"`
	HashPassword string `json:"hash_password" gorm:"column:hash_password;" binding:"required"`
	Age          string `json:"age" gorm:"column:age;" binding:"required"`
	Gender       string `json:"gender" gorm:"column:gender;"`
}

type Card struct {
	gorm.Model
	Email string `gorm:"column:email"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create ... create user
func (u *UserModel) CreateUser() error {
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Save ... save user.
func (u *UserModel) Save() error {
	return model.DB.Self.Save(u).Error
}

// Get Information
func (user *UserModel) GetInfo(email string) error {
	fmt.Printf("email:%s\n", email)
	if err := model.DB.Self.Model(UserModel{}).
		Where("email = ?", email).
		First(user).Error; err != nil {
		fmt.Println("error", err)
		return err
	}
	return nil
}

func IfExist(email, name string) error {
	var user1 UserModel
	var user2 UserModel

	err1 := model.DB.Self.Debug().Where("email=?", email).Find(&user1).Error
	err2 := model.DB.Self.Debug().Where("name=?", name).Find(&user2).Error

	s := []string{""}
	i := 0

	if err1 == nil {
		s = append(s, "*邮箱*")
		i++
	}

	if err2 == nil {
		s = append(s, "*姓名*")
		i++

	}

	if i > 0 {
		s = append(s, "已被注册")
	}

	if i > 0 {
		return errors.New(fmt.Sprintf("%s", s))
	}

	return nil

}

func UpdateInfo(email string, avatar string, name string) error {
	var user = UserModel{
		Email:  email,
		Name:   name,
		Avatar: avatar,
	}
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(user).Where("email = ?", user.Email).Update(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateInfor(email string, avatar string, name string, studentId string, college string, major string, grade string, gender string, phone_number string, qq_number string) error {
	var user = UserModel{
		Email:  email,
		Avatar: avatar,
		Name:   name,
		Gender: gender,
	}
	tx := model.DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(user).Where("email = ?", user.Email).Update(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdatePassword(email string, original string, new string) error {
	md5 := Md5.New()
	md5.Write([]byte(original))
	hashPwd := hex.EncodeToString(md5.Sum(nil))

	var user UserModel
	if err := model.DB.Self.Model(UserModel{}).Where("email = ?", email).Find(&user).Error; err != nil {
		return err
	}
	if hashPwd != user.HashPassword {
		return errors.New("original password error")
	}
	newmd5 := Md5.New()
	newmd5.Write([]byte(new))
	newhashPwd := hex.EncodeToString(md5.Sum(nil))
	if err := model.DB.Self.Model(UserModel{}).Where("email = ?", email).Update(UserModel{HashPassword: newhashPwd}).Error; err != nil {
		return err
	}
	return nil
}

func PunchCard(email string) error {
	var card Card

	if err := model.DB.Self.Model(&Card{}).Where("email = ?", email).First(&card).Error; err != nil {
		var new_card = Card{
			Email: email,
		}
		if err := model.DB.Self.Model(&Card{}).Create(&new_card).Error; err != nil {
			return err
		} else {
			return nil
		}
	}

	y1, m1, d1 := card.Model.CreatedAt.Date()
	y2, m2, d2 := time.Now().Date()
	if y1 != y2 || m1 != m2 || d1 != d2 {
		var new_card = Card{
			Email: email,
		}
		if err := model.DB.Self.Model(&Card{}).Create(&new_card).Error; err != nil {
			return err
		}
	} else {
		errors.New("have punched card today")
	}
	return nil
}

func GetCard(email, year, month string) ([]*Card, error) {
	cards := make([]*Card, 0)
	var where string
	switch month {
	case "1", "3", "5", "7", "8", "10", "12":
		where = fmt.Sprintf("email = '%v' AND created_at >= '%v-%v-01 00:00:00' AND created_at <= '%v-%v-31 23:59:59'", email, year, month, year, month)
	case "4", "6", "9", "11":
		where = fmt.Sprintf("email = '%v' AND created_at >= '%v-%v-01 00:00:00' AND created_at <= '%v-%v-30 23:59:59'", email, year, month, year, month)
	default:
		where = fmt.Sprintf("email = '%v' AND created_at >= '%v-%v-01 00:00:00' AND created_at <= '%v-%v-28 23:59:59'", email, year, month, year, month)

	}
	if err := model.DB.Self.Model(&Card{}).Where(where).Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}
