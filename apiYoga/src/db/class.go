package db

func InsertNewClass(classList ClassList) (classId int, err error) {
	err = postdb.Model(&ClassList{}).Create(&classList).Error
	return classId, err
}
func DeleteClass(classId int) (err error) {
	err = postdb.Where("class_id =?", classId).Delete(&ClassList{}).Error
	return err
}
