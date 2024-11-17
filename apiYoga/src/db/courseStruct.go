package db

import "time"

type Course struct {
	CourseID        int `gorm:"primaryKey"`
	AdminID         int
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
