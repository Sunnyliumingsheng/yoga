package db

import (
	"gorm.io/gorm"

	"api/config"
)

// 检查用户表中是否已经存在这个openid的用户
func IsThisOpenIdExistedAndGetLevel(openid string) (isExist bool, userIdIfExist int, level int, err error) {
	var user User
	err = postdb.Model(&User{}).Where("openid = ?", openid).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, -1, -1, nil
	}
	if err != nil {
		return false, -1, -1, err
	}
	level = 4
	if user.IsStudent {
		level = 3
	}
	if user.IsTeacher {
		level = 2
	}
	if user.IsAdmin {
		level = 1
	}
	return true, int(user.UserID), level, nil
}
func InsertUserAndGetUserId(openid string) (userId int, err error) {
	user := User{
		OpenID:    openid,
		Nickname:  config.Config.NewUserDefaultInfo.Nickname,
		Gender:    config.Config.NewUserDefaultInfo.Gender,
		Signature: config.Config.NewUserDefaultInfo.Signature,
		AvaURL:    config.Config.NewUserDefaultInfo.AvaURL,
		IsStudent: false,
		IsTeacher: false,
		IsAdmin:   false,
	}
	err = postdb.Create(&user).Error
	if err != nil {
		return -1, err
	}
	return user.UserID, nil
}
