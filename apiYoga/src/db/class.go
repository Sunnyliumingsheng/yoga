package db

func InsertNewClass(classList ClassList) (classId int, err error) {
	err = postdb.Model(&ClassList{}).Create(&classList).Error
	return classId, err
}
func DeleteClass(classId int) (err error) {
	err = postdb.Where("class_id =?", classId).Delete(&ClassList{}).Error
	return err
}
func ActiveClass(classId int) (err error) {
	err = postdb.Model(&ClassList{}).Where("class_id =?", classId).Update("already_active", true).Error
	return err
}
func SelectActivedClassThisWeekday(weekday int) (classList []ClassList, err error) {
	err = postdb.Model(&ClassList{}).Where("day_of_week =? AND already_active =?", weekday, true).Find(&classList).Error
	return classList, err
}
func SelectAllActivedClass() (classList []ClassList, err error) {
	err = postdb.Model(&ClassList{}).Where("already_active =?", true).Find(&classList).Error
	return classList, err
}
func SelectAllClass() (classList []ClassList, err error) {
	err = postdb.Find(&classList).Error
	return classList, err
}
