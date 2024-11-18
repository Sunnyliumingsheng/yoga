package db

import "time"

type Course struct {
	CourseId        int `gorm:"primaryKey"`
	AdminId         int
	CourseName      string `gorm:"size:255"`
	CourseSubject   string
	IntroductionURL string `gorm:"size:255"`
	Introduction    string
	IsGroupType     bool
	IsTeamType      bool
	IsVipType       bool
	RecommendMaxNum int
	RecommendMinNum int
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
type CourseBasic struct {
	CourseID   int    `json:"courseID"`
	CourseName string `json:"courseName"`
}
