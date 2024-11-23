package deamon

import (
	"api/db"
	"api/loger"
	"errors"
	"time"
)

var ccb [4]db.OneDayClass
var nowWeekDay int
var pmap int

// when program start to do,init the data structure
func InitActivedClass() {
	weekday := time.Now().Weekday()
	nowWeekDay = int(weekday)
	for i := 0; i <= 3; i++ {
		classList, err := db.SelectActivedClassThisWeekday((nowWeekDay + i) % 7)
		if err != nil {
			loger.Loger.Println("error:", err.Error())
		}
		var classActived map[int]db.ClassActived
		classActived = make(map[int]db.ClassActived)
		var userResumeInfo map[int][]db.UserResumeInfo
		userResumeInfo = make(map[int][]db.UserResumeInfo)
		for _, class := range classList {
			classActived[class.ClassId] = db.ClassActived{
				CourseId:  class.CourseId,
				Index:     class.Index,
				ResumeNum: 0,
				TeacherId: class.TeacherId,
				Max:       class.Max,
			}
			userResumeInfo[class.ClassId] = make([]db.UserResumeInfo, 0, class.Max)
		}
		ccb[i] = db.OneDayClass{
			ClassActived:   classActived,
			UserResumeInfo: userResumeInfo,
		}
	}
	pmap = 0
}

// renew the data structure,every night todo
func RenewActivedClass() {
	// firstly record the resume information
	for classId, classInfo := range ccb[pmap].ClassActived {
		recordId, err := db.InsertClassRecord(classId, classInfo)
		if err != nil {
			loger.Loger.Println("error: insert class record failed", err.Error())
			return
		}
		err = db.InsertCheckinRecord(recordId, ccb[pmap].UserResumeInfo[classId])
		if err != nil {
			loger.Loger.Println("error : insert checkin record failed", err.Error())
			return
		}
	}
	// then delete the old map
	ccb[pmap] = db.OneDayClass{}
	nowWeekDay = ((nowWeekDay + 1) % 7)
	newActivedClassList, err := db.SelectActivedClassThisWeekday(nowWeekDay)
	if err != nil {
		loger.Loger.Println("error:", err.Error())
		return
	}
	for _, class := range newActivedClassList {
		ccb[pmap].ClassActived[class.ClassId] = db.ClassActived{
			CourseId:   class.CourseId,
			Index:      class.Index,
			ResumeNum:  0,
			CheckinNum: 0,
			TeacherId:  class.TeacherId,
			Max:        class.Max,
			RecordText: "",
		}
		ccb[pmap].UserResumeInfo[class.ClassId] = make([]db.UserResumeInfo, 0, class.Max)
	}
	pmap = (pmap + 1) % 4
}

func QuicklySelectClass() (fourDayClass [4][]db.ActivedClassInfo, err error) {
	for i := 0; i <= 3; i++ {
		for classId, classActived := range ccb[i].ClassActived {
			fourDayClass[i] = append(fourDayClass[i], db.ActivedClassInfo{
				ClassId:      classId,
				ClassActived: classActived,
			})
		}
	}
	return fourDayClass, nil
}

// give the class basic info and all the resume info
func QuicklySelectClassByClassId(classId int) (db.ClassActived, []db.UserResumeInfo, error) {
	for i := 0; i <= 3; i++ {
		// 很典型的奇技淫巧，很容易出错，切勿模仿
		if classActived, ok := ccb[i].ClassActived[classId]; ok {
			return classActived, ccb[i].UserResumeInfo[classId], nil
		}
	}
	return db.ClassActived{}, nil, errors.New("找不到这个课程")
}

// student try to resume
func Resume(userId, classId int) (err error) {
	// find the class in this 4 day
	for i := 0; i <= 3; i++ {
		classActived, ok := ccb[i].ClassActived[classId]
		if ok { //if i find the class
			// select the resume info
			userResumeInfo, ok := ccb[i].UserResumeInfo[classId]
			if !ok {
				loger.Loger.Println("error: can find the class info but not the resume info", classId, userId)
				return errors.New("出现错误，通知管理员")
			}
			// check the user not repeating the resume
			for _, value := range userResumeInfo {
				if value.UserId == userId {
					return errors.New("已经预约过了,请不要重复预约")
				}
			}
			basicCardInfo, ok := UserCard[userId]
			if !ok {
				loger.Loger.Println("error: can find the card info", userId, classId)
				return errors.New("出现错误，可以联系管理员")
			}
			err := db.UpdateCardIfTimesCardByUserId(userId, basicCardInfo.PurchaseId)
			if err != nil {
				loger.Loger.Println("error: 卡次数更新失败 userId:", userId, "error:", err.Error())
				return err
			}
			userResumeInfo = append(userResumeInfo, db.UserResumeInfo{
				UserId: userId,
				Status: 0,
			})
			classActived.ResumeNum++
			return nil
		}
	}
	loger.Loger.Println("error: cant find this class may encounter attack", classId)
	return errors.New("找不到这个课程，请联系管理员")
}
func QuicklySelectMyResume(userId int) (activedClassInfo db.ActivedClassInfo, userResumeInfo []db.UserResumeInfo, err error) {
	for i := 0; i <= 3; i++ {

	}
}
