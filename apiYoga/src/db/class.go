package db

import (
	"api/config"
	"time"

	"gorm.io/gorm"
)

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
func SelectCourseIdByClassId(classId int) (courseId int, err error) {
	var class ClassList
	err = postdb.Where("class_id =?", classId).First(&class).Error
	if err != nil {
		return 0, err
	}
	return class.CourseId, err
}
func SelectWeekdayByClassId(classId int) (dayOfWeek int, err error) {
	var class ClassList
	err = postdb.Where("class_id =?", classId).First(&class).Error
	if err != nil {
		return 0, err
	}
	return class.DayOfWeek, err
}
func SelectTeachClassThisWeekday(teacherId, weekday int) (class []ClassList, err error) {
	err = postdb.Model(&ClassList{}).Where("teacher_id=? AND day_of_week=?", teacherId, weekday).Find(&class).Error
	if err != nil {
		return nil, err
	}
	return class, nil
}
func InsertClassRecord(classId int, classInfo ClassActived) (classRecordId int, err error) {
	var class ClassRecord
	class = ClassRecord{
		ClassId:       classId,
		EndTime:       time.Now(),
		ShouldCheckin: classInfo.ResumeNum,
		ReallyCheckin: classInfo.CheckinNum,
		RecordText:    classInfo.RecordText,
	}
	err = postdb.Model(&ClassRecord{}).Create(&class).Error
	return class.ClassRecordId, err
}
func InsertCheckinRecord(classRecordId int, resumeInfos []UserResumeInfo) (err error) {
	var checkInRecords []CheckinRecord = make([]CheckinRecord, 20)
	for index, resumeInfo := range resumeInfos {
		checkInRecords[index] = CheckinRecord{
			ClassRecordId: classRecordId,
			UserId:        resumeInfo.UserId,
			Status:        resumeInfo.Status,
			CheckinAt:     resumeInfo.CheckinAt,
		}
	}
	err = postdb.Model(&CheckinRecord{}).Create(checkInRecords).Error
	return err
}
func StorageClassRam(ramInfo StorageStruct) {
	postdb.Model(&StorageStruct{}).Unscoped().Where("1 = 1").Delete(&StorageStruct{})
	postdb.Model(&StorageStruct{}).Create(ramInfo)
}
func GetSTorageClassRam() (ramInfo StorageStruct) {
	postdb.Model(&StorageStruct{}).First(&ramInfo)
	return ramInfo
}
func SelectRecord(tail int) ([]ClassRecord, error) {
	var classRecords []ClassRecord = make([]ClassRecord, tail)
	err := postdb.Model(&ClassRecord{}).Order("end_time DESC").Limit(tail).Find(&classRecords).Error
	if err != nil {
		return nil, err
	}
	return classRecords, nil
}

// night record those checkin and class message
func RecordClassAndCheckinInfo(oneDayInfo OneDayClass) error {
	var classRecords []ClassRecord
	classRecords = make([]ClassRecord, 4)
	var checkinRecords []CheckinRecord
	checkinRecords = make([]CheckinRecord, 50)
	for classId, classInfo := range oneDayInfo.ClassActivedBlock {
		classRecords = append(classRecords, ClassRecord{
			ClassId:       classId,
			EndTime:       time.Now(),
			ShouldCheckin: classInfo.ResumeNum,
			ReallyCheckin: classInfo.CheckinNum,
			RecordText:    classInfo.RecordText,
			Index:         classInfo.Index,
			CourseId:      classInfo.CourseId,
			TeacherId:     classInfo.TeacherId,
			WeekDay:       classInfo.WeekDay,
		})
	}
	err := postdb.Model(&ClassRecord{}).Save(&classRecords).Error
	if err != nil {
		return err
	}
	for _, classRecord := range classRecords {
		for _, resumeInfo := range oneDayInfo.UserResumeInfoBlock[classRecord.ClassId] {
			checkinRecords = append(checkinRecords, CheckinRecord{
				ClassRecordId: classRecord.ClassRecordId,
				UserId:        resumeInfo.UserId,
				Status:        resumeInfo.Status,
				CheckinAt:     resumeInfo.CheckinAt,
			})
		}
		err := postdb.Model(&CheckinRecord{}).Save(&checkinRecords).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func InsertBlackList(userId int) (err error) {
	err = postdb.Model(&Blacklist{}).Create(Blacklist{
		UserId: userId,
		EndAt:  time.Now().AddDate(0, 0, config.Config.Rules.BlackListTime),
	}).Error
	return err
}
func SelectBlackList() (blackLists []Blacklist, err error) {
	err = postdb.Find(&blackLists).Error
	return blackLists, err
}
func IsUserIdNotInBlackList(userId int) (bool, error) {
	var blacklist Blacklist
	err := postdb.Where("user_id =?", userId).First(&blacklist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, err
	}
	if blacklist.EndAt.After(time.Now()) {
		return false, nil
	}
	return true, nil
}

// 可以经常使用
func DeleteAllBlackList() (err error) {
	err = postdb.Delete(&Blacklist{}).Error
	return err
}

// 配合这个函数，可以使用上面这个
func IsSomeoneInBlackList(userId int) (bool, error) {
	var blacklist Blacklist
	err := postdb.Where("user_id =?", userId).First(&blacklist).Error
	if err == nil {
		return true, nil
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, err
}
