package db

import "gorm.io/gorm"

func insertUserByNameCheckNameIsExistEngine(name string) bool {
	var user User
	err := postdb.Model(&User{}).Where("name=?", name).First(&user).Error
	return !(err == gorm.ErrRecordNotFound)
}
