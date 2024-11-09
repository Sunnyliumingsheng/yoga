package db

import (
	"errors"
	"fmt"

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

func HandleUserLevelStudent(errChan chan error, user *User, wantIsStudent bool) {
	if user.IsStudent {
		if wantIsStudent {
			//已经是一个学生了,并且想要改为一个学生
			errChan <- nil
			return
		} else {
			//目前是一个学生,想要取消他的学生身份
			user.IsStudent = false
			err := postdb.Model(&User{}).Where("user_id=?", user.UserID).Update("is_student", false).Error
			if err != nil {
				errChan <- err
				return
			}
			err = dropStudentByUserId(user.UserID)
			if err != nil {
				errChan <- err
				return
			}
			errChan <- nil
			return
		}
	} else {
		if wantIsStudent {
			// 目前不是一个学生, 并且想要将他拥有学生身份
			user.IsStudent = true
			err := postdb.Save(user).Error
			if err != nil {
				errChan <- err
				return
			}
			err = insertNewStudent(user.UserID)
			if err != nil {
				errChan <- err
				return
			}
			errChan <- nil
			return
		} else {
			// 目前不是一个学生, 并且不想要将他变为非学生身份
			errChan <- nil
			return
		}
	}
}
func HandleUserLevelTeacher(errChan chan error, user *User, wantIsTeacher bool) {
	if user.IsTeacher {
		if wantIsTeacher {
			// 已经是一个老师了, 并且想要改为一个老师
			errChan <- nil
			return
		} else {
			// 目前是一个老师, 并且想要将他的老师身份取消
			user.IsTeacher = false
			err := postdb.Model(&User{}).Where("user_id=?", user.UserID).Update("is_teacher", false).Error
			if err != nil {
				errChan <- err
				return
			}
			err = dropTeacherByUserId(user.UserID)
			if err != nil {
				errChan <- err
				return
			}
			errChan <- nil
			return
		}
	} else {
		if wantIsTeacher {
			// 目前不是一个老师, 并且想要将他拥有老师身份
			user.IsTeacher = true
			err := postdb.Save(user).Error
			if err != nil {
				errChan <- err
				return
			}
			err = insertNewTeacher(user.UserID)
			if err != nil {
				errChan <- err
				return
			}
			errChan <- nil
			return
		} else {
			// 目前不是一个老师, 并且不想要将他变为非老师身份
			errChan <- nil
			return
		}
	}
}
func HandleUserLevelAdmin(errChan chan error, user *User, wantIsAdmin bool) {
	if user.IsAdmin {
		if wantIsAdmin {
			// 已经是一个管理员了, 并且想要改为一个管理员
			errChan <- nil
			return
		} else {
			// 目前是一个管理员, 并且想要将他的管理员身份取消
			user.IsAdmin = false
			err := postdb.Model(&User{}).Where("user_id=?", user.UserID).Update("is_admin", false).Error
			if err != nil {
				errChan <- err
				return
			}
			err = dropAdminByUserId(user.UserID)
			if err != nil {
				errChan <- err
				return
			}
			errChan <- nil
			return
		}
	} else {
		if wantIsAdmin {
			// 目前不是一个管理员, 并且想要将他拥有管理员身份
			user.IsAdmin = true
			err := postdb.Save(user).Error
			if err != nil {
				errChan <- err
				return
			}
			err = insertNewAdmin(user.UserID)
			if err != nil {
				errChan <- err
				return
			}
			errChan <- nil
			return
		} else {
			// 目前不是一个管理员, 并且不想要将他变为非管理员身份
			errChan <- nil
			return
		}
	}
}
func Rename(userId string, newName string) (err error) {
	var user User
	err = postdb.Model(&User{}).Where("name=?", newName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		err = postdb.Model(&User{}).Where("user_id=?", userId).Update("name", newName).Error
		return err
	} else {
		if err == nil {
			return errors.New("存在重名,请重新命名,推荐使用真名+数字再尝试")
		} else {
			return errors.New("意料之外的错误,请联系管理员" + err.Error())
		}
	}

}

// 就像linux中的tail一样,选择性的查看几条用户
func SelectUserTail(tail int) (usersInfo []User, err error) {
	err = postdb.Model(&User{}).Limit(tail).Find(&usersInfo).Error
	if err != nil {
		return nil, err
	}
	return usersInfo, nil
}
func SelectAdminInfo(userId int) (adminInfo Admin, err error) {
	adminInfo.UserID = userId
	err = postdb.First(&adminInfo).Error
	return adminInfo, err
}
func SelectTeacherInfo(userId int) (teacherInfo Teacher, err error) {
	teacherInfo.UserID = userId
	err = postdb.First(&teacherInfo).Error
	return teacherInfo, err
}
func SelectAndCheckUserInfoByName(name string) (user User, isExist bool, err error) {
	err = postdb.Model(&User{}).Where("name = ?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return User{}, false, nil
	}
	if err != nil {
		return User{}, false, err
	}
	return user, true, nil
}
func InsertAdminAccountAndPassword(adminId int, account, password string) (isExist bool, err error) {
	//检查是否已经存在这个admin的account
	var admin Admin
	admin.Account = account
	err = postdb.Model(&Admin{}).Where("account=?", account).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if err == nil {
		return true, nil
	}

	err = postdb.Model(&Admin{}).Where("admin_id = ?", adminId).Updates(map[string]interface{}{"account": account, "password": password}).Error
	return false, err
}

func InsertTeacherAccountAndPassword(teacherId int, account, password string) (isExist bool, err error) {
	var teacher Teacher
	teacher.Account = account
	err = postdb.Model(&Teacher{}).Where("account=?", account).First(&teacher).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("here is error:")
		return false, err
	}
	if err == nil {
		return true, nil
	}
	err = postdb.Model(&Teacher{}).Where("teacher_id = ?", teacherId).Updates(map[string]interface{}{"account": account, "password": password}).Error
	return false, err

}

func TeacherLogin(account, password string) (isOk bool, err error) {
	var teacher Teacher
	err = postdb.Model(&Teacher{}).Where("account=? AND password=?", account, password).First(&teacher).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, nil
	}
	return true, err
}

func AdminLogin(account, password string) (isOk bool, err error) {
	var admin Admin
	fmt.Println("get is:", account, password)
	err = postdb.Model(&Admin{}).Where("account=? AND password=?", account, password).First(&admin).Error
	fmt.Println("error:", err, "data:", admin)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, nil
	}
	return true, err
}
func SelectAdminIdByAdminAccount(account string) (adminId int, err error) {
	var admin Admin
	err = postdb.Model(&Admin{}).Where("account=?", account).First(&admin).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return admin.AdminID, err
}
func SelectAdminInfoByAccount(account string) (adminInfo Admin, err error) {
	err = postdb.Model(&adminInfo).Where("account=?", account).First(&adminInfo).Error
	return adminInfo, err
}
func SelectUserInfoByUserId(userId int) (userInfo User, err error) {
	err = postdb.Model(&userInfo).Where("user_id=?", userId).First(&userInfo).Error
	return userInfo, err
}
func SelectTeacherInfoByAccount(account string) (teacherInfo Teacher, err error) {
	err = postdb.Model(&teacherInfo).Where("account=?", account).First(&teacherInfo).Error
	return teacherInfo, err
}

// 警惕 ，来自gpt，非手写
func UpdateUserInfo(userId, nickname, signature string, gender bool) (err error) {
	updateData := map[string]interface{}{
		"nickname":  nickname,
		"signature": signature,
		"gender":    gender,
	}

	// 使用 gorm 的 Updates 方法进行更新
	err = postdb.Model(&User{}).Where("user_id = ?", userId).Updates(updateData).Error
	return
}
func UpdateTeacherInfo(userId, introduction string) (err error) {
	updateData := map[string]interface{}{
		"introduction": introduction,
	}
	postdb.Model(&Teacher{}).Where("user_id =?", userId).Updates(updateData)
	return err
}
