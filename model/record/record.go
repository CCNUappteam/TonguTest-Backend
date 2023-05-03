package record

import (
	"github.com/jinzhu/gorm"
	"time"
	"tongue/model"
)

type Record struct {
	gorm.Model
	Email  string `json:"email" gorm:"column:email"`
	Health string `json:"health" gorm:"column:health"`
}

func CreateRecord(email string, health string) error {
	var record Record

	err := model.DB.Self.Model(&Record{}).Where("email = ?").First(&record).Error
	if err != nil {
		record = Record{
			Email:  email,
			Health: health,
		}
		if err := model.DB.Self.Model(&Record{}).Create(&record).Error; err != nil {
			return err
		}
		return nil
	}
	if record.CreatedAt == time.Now() {
		record.Health = health
		if err := model.DB.Self.Model(&Record{}).Update(&record).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetRecord(email string) ([]*Record, error) {
	var record = make([]*Record, 0)
	if err := model.DB.Self.Model(&Record{}).Where("email = ?", email).Find(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}
