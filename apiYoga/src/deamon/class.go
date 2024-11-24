package deamon

import (
	"api/db"
	"api/loger"
	"errors"
	"fmt"
	"time"
)

var ccb [4]db.OneDayClass
var nowWeekDay int
var pmap int

// used to debug,print the ccb information now
func Print() {
	for _, oneCcb := range ccb {
		fmt.Println("--------------------------")
		fmt.Println(oneCcb.ClassActivedBlock)
		fmt.Println("-----")
		fmt.Println(oneCcb.ClassActivedBlock)
		fmt.Println("--------------------------")
	}
}

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
			ClassActivedBlock:   classActived,
			UserResumeInfoBlock: userResumeInfo,
		}
	}
	pmap = 0
}

// renew the data structure,every night todo
func RenewActivedClass() {
	// firstly record the resume information
	for classId, classInfo := range ccb[pmap].ClassActivedBlock {
		recordId, err := db.InsertClassRecord(classId, classInfo)
		if err != nil {
			loger.Loger.Println("error: insert class record failed", err.Error())
			return
		}
		err = db.InsertCheckinRecord(recordId, ccb[pmap].UserResumeInfoBlock[classId])
		if err != nil {
			loger.Loger.Println("error : insert checkin record failed", err.Error())
			return
		}
	}
	// then delete the old map
	err := db.RecordClassAndCheckinInfo(ccb[pmap])
	if err != nil {
		loger.Loger.Println("error : when record this day ", err)
		return
	}
	ccb[pmap] = db.OneDayClass{}
	// and insert the new class
	nowWeekDay = ((nowWeekDay + 1) % 7)
	newActivedClassList, err := db.SelectActivedClassThisWeekday(nowWeekDay)
	if err != nil {
		loger.Loger.Println("error:", err.Error())
		return
	}
	for _, class := range newActivedClassList {
		ccb[pmap].ClassActivedBlock[class.ClassId] = db.ClassActived{
			CourseId:   class.CourseId,
			Index:      class.Index,
			ResumeNum:  0,
			CheckinNum: 0,
			TeacherId:  class.TeacherId,
			Max:        class.Max,
			WeekDay:    class.DayOfWeek,
			RecordText: "",
		}
		ccb[pmap].UserResumeInfoBlock[class.ClassId] = make([]db.UserResumeInfo, 0, class.Max)
	}
	pmap = (pmap + 1) % 4
}

func QuicklySelectClass() (fourDayClass [4][]db.ActivedClassInfo, err error) {
	for i := 0; i <= 3; i++ {
		for classId, classActived := range ccb[i].ClassActivedBlock {
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
		if classActived, ok := ccb[i].ClassActivedBlock[classId]; ok {
			return classActived, ccb[i].UserResumeInfoBlock[classId], nil
		}
	}
	return db.ClassActived{}, nil, errors.New("找不到这个课程")
}

// student try to resume
func Resume(userId, classId int) (err error) {
	// find the class in this 4 day
	for i := 0; i <= 3; i++ {
		classActived, ok := ccb[i].ClassActivedBlock[classId]
		if ok { //if i find the class
			// select the resume info
			userResumeInfo, ok := ccb[i].UserResumeInfoBlock[classId]
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

// select my resume
func QuicklySelectMyResume(userId int) (resumeInfo db.OneDayClass) {
	MyResumeClassId := make([]int, 0, 8)
	for i := 0; i <= 3; i++ {
		for classId, ResumeInfo := range ccb[i].UserResumeInfoBlock {
			for _, value := range ResumeInfo {
				// find all classId that this userId resumed
				if value.UserId == userId {
					MyResumeClassId = append(MyResumeClassId, classId)
				}
			}
		}
		for _, classId := range MyResumeClassId {
			resumeInfo.ClassActivedBlock[classId] = ccb[i].ClassActivedBlock[classId]
			resumeInfo.UserResumeInfoBlock[classId] = ccb[i].UserResumeInfoBlock[classId]
		}
	}
	return resumeInfo
}

func QuicklyCancelResume(userId, classId int) (isok bool, err error) {
	weekday, resumeInfo, err := lockResumeInfoByClassIdAndUserId(classId, userId)
	if nowWeekDay == weekday {

		return false, nil
	}
	resumeInfo.Status = 3
	return true, err
}

// this function may error ,because it too relay human brain,return weekday and the pointer of resume struct
func lockResumeInfoByClassIdAndUserId(classId, userId int) (int, *db.UserResumeInfo, error) {
	for i := 0; i <= 3; i++ {
		resumeInfoBlock, ok := ccb[i].UserResumeInfoBlock[classId]
		if !ok { // cant find this class
			continue
		}
		// can find this class and range this resume block
		for _, resumeInfo := range resumeInfoBlock {
			if resumeInfo.UserId == userId {
				return ccb[i].ClassActivedBlock[classId].WeekDay, &resumeInfo, nil
			}
		}
		loger.Loger.Println("找不到预约记录", userId)
		return -1, nil, errors.New("找不到预约记录")
	}
	loger.Loger.Println("找不到这个班级", classId)
	return -1, nil, errors.New("找不到这个班级")
}

// thi function lock the classId
func lockResumeInfoByClassId(classId int) (i int, err error) {
	for i = 0; i <= 3; i++ {
		_, ok := ccb[i].UserResumeInfoBlock[classId]
		if !ok { // cant find this class
			continue
		}
		// find this class
		return i, nil
	}
	return -1, errors.New("cant find the class")
}
func StorageClass() {
	Print()
	db.StorageClassRam(db.StorageStruct{
		Ccb:        ccb,
		NowWeekDay: nowWeekDay,
		Pmap:       pmap,
	})
}
func GetStorageClass() {
	storage := db.GetSTorageClassRam()
	ccb = storage.Ccb
	nowWeekDay = storage.NowWeekDay
	pmap = storage.Pmap
	Print()
}
func CheckinAllStudent(userId, classId int, text string) (err error) {
	p, err := lockResumeInfoByClassId(classId)
	if err != nil {
		return err
	}
	classInfo, ok := ccb[p].ClassActivedBlock[classId]
	if !ok {
		return errors.New("出现问题，无法找到这节课")
	}
	if classInfo.TeacherId != userId {
		return errors.New("你不是这门课程的老师")
	}
	classInfo.RecordText = text
	for _, resumeInfo := range ccb[p].UserResumeInfoBlock[classId] {
		resumeInfo.Status = 1
		resumeInfo.CheckinAt = time.Now()
	}
	return nil
}
func ChangeCheckinStatusUser(userId, status, classId int) (err error) {
	_, resume, err := lockResumeInfoByClassIdAndUserId(classId, userId)
	if err != nil {
		return err
	}
	resume.Status = status
	// you cant return err ,because may change status
	return nil
}
