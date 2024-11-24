package db

import "time"

type ClassList struct {
	ClassId       int `gorm:"primaryKey;autoIncrement"`
	ClassName     string
	CourseId      int
	Auto          bool
	DayOfWeek     int
	AlreadyActive bool
	Index         int
	Min           int
	Max           int
	TeacherId     int
}

type ClassRecord struct {
	ClassRecordId int `gorm:"primaryKey;autoIncrement"`
	ClassId       int
	EndTime       time.Time
	ShouldCheckin int
	ReallyCheckin int
	Index         int
	CourseId      int
	TeacherId     int
	WeekDay       int
	RecordText    string
}

type CheckinRecord struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	ClassRecordId int
	UserId        int
	Status        int
	CheckinAt     time.Time
}

type Blacklist struct {
	ID     int `gorm:"primaryKey;autoIncrement"`
	UserId int
	EndAt  time.Time
}

type StorageStruct struct {
	ID         int `gorm:"primaryKey"`
	Ccb        [4]OneDayClass
	NowWeekDay int
	Pmap       int
}

// 这几个存在于内存中，每天凌晨和重启的时候才会用到
type ClassActived struct {
	Index      int
	CourseId   int
	ResumeNum  int
	CheckinNum int
	TeacherId  int
	Max        int
	WeekDay    int
	RecordText string
}
type UserResumeInfo struct {
	UserId    int
	Status    int
	CheckinAt time.Time
}
type OneDayClass struct {
	// the int is the classId
	ClassActivedBlock   map[int]ClassActived
	UserResumeInfoBlock map[int][]UserResumeInfo
}

// this struct is used to response the request
type ActivedClassInfo struct {
	ClassId      int
	ClassActived ClassActived
}
