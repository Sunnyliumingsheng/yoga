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
	StartTime     time.Time
	EndTime       time.Time
	ShouldCheckin int
	ReallyCheckin int
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

// 这个存在于内存中，每天凌晨和重启的时候才会用到
type ClassActived struct {
	ClassId   int
	Index     int
	ResumeNum int
	TeacherId int
}
type UserResumeInfo struct {
	userId    int
	status    int
	checkinAt time.Time
}
