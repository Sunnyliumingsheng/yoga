package db

import "gorm.io/gorm"

func insertUserByNameCheckNameIsExistEngine(name string) bool {
	var user User
	err := postdb.Model(&User{}).Where("name=?", name).First(&user).Error
	return !(err == gorm.ErrRecordNotFound)
}
func insertNewStudent(userId int) (err error) {
	student := Student{
		UserID: userId,
	}
	err = postdb.Create(&student).Error
	return err
}
func insertNewTeacher(userId int) (err error) {
	teacher := Teacher{
		UserID: userId,
	}
	err = postdb.Create(&teacher).Error
	return err
}

func insertNewAdmin(userId int) (err error) {
	admin := Admin{
		UserID: userId,
	}
	err = postdb.Create(&admin).Error
	return err
}

func dropStudentByUserId(userId int) (err error) {
	err = postdb.Where("user_id = ?", userId).Delete(&Student{}).Error
	return err
}
func dropTeacherByUserId(userId int) (err error) {
	err = postdb.Where("user_id = ?", userId).Delete(&Teacher{}).Error
	return err
}
func dropAdminByUserId(userId int) (err error) {
	err = postdb.Where("user_id = ?", userId).Delete(&Admin{}).Error
	return err
}

// 这个函数可能暂时都用不上
func SelectUserInfoFromStuTeaAdm(userId int) (student []Student, teacher []Teacher, admin []Admin, err error) {

	selectChan := make(chan error, 3)
	go func() {
		selectChan <- postdb.Where("user_id =?", userId).Find(&student).Error
	}()
	go func() {
		selectChan <- postdb.Where("user_id =?", userId).Find(&teacher).Error
	}()
	go func() {
		selectChan <- postdb.Where("user_id =?", userId).Find(&admin).Error
	}()
	for i := 0; i < 3; i++ {
		if err := <-selectChan; err != nil && err != gorm.ErrRecordNotFound {
			return nil, nil, nil, err
		}
	}
	close(selectChan)
	return student, teacher, admin, nil
}
