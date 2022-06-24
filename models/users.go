package models

import (
	"gorm.io/gorm"
)

type UsersValues struct {
	//gorm.Model
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   int64  `gorm:"default:1" json:"status"`
}

func CreateUsers(db *gorm.DB, users *UsersValues) (err error) {
	err = db.Create(users).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB, users *[]UsersValues) (err error) {
	err = db.Find(users).Error
	if err != nil {
		return err
	}
	return nil
}
