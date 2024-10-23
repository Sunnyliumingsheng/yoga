package db

import (
	"errors"

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
		Openid:    openid,
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
func IntUserIdSelectUserLevel(userId int) (level int, err error) {
	var user User
	err = postdb.Model(&User{}).Where("user_id =?", userId).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return -1, gorm.ErrRecordNotFound
	}
	if err != nil {
		return -1, err
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
	return level, nil
}
func SelectUserInfoByName(name string) (user User, err error) {
	err = postdb.Model(&User{}).Where("name = ?", name).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func InsertNewUser(name string) (ExistName bool, err error) {
	if insertUserByNameCheckNameIsExistEngine(name) {
		return true, errors.New("已经存在这个名称的用户了")
	}
	user := User{
		Name:      name,
		Nickname:  config.Config.NewUserDefaultInfo.Nickname,
		Gender:    config.Config.NewUserDefaultInfo.Gender,
		Signature: config.Config.NewUserDefaultInfo.Signature,
		AvaURL:    config.Config.NewUserDefaultInfo.AvaURL,
		IsStudent: false,
		IsTeacher: false,
		IsAdmin:   false,
	}
	result := postdb.Model(&User{}).Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return false, nil

}
func DropUserByStringUserId(userId string) (err error) {
	err = postdb.Delete(&User{}, userId).Error
	if err != nil {
		return err
	}
	return nil
}
