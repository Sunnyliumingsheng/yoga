package db

import "time"

// 存储在数据库中
// 这里是所有卡的列表，由管理员创建，用户从这里检索和挑选，系统从这里获取卡片信息
type CardList struct {
	CardId               int    `gorm:"primaryKey;autoIncrement"` // SERIAL PRIMARY KEY
	CardName             string `gorm:"unique;size:255"`          // VARCHAR(255) UNIQUE
	BriefIntroductionURL string `gorm:"size:255"`                 // VARCHAR(255)
	AdminAccount         string `gorm:"size:255"`                 // VARCHAR(255)
	CardIntroduction     string `gorm:"type:text"`                // TEXT
	CardIntroductionURL  string `gorm:"size:255"`                 // VARCHAR(255)
	IsSupportGroup       bool
	IsSupportTeam        bool
	IsSupportVIP         bool
	IsLimitDays          bool
	IsLimitTimes         bool
	IsForbidSpecial      bool
	IsSupportSpecial     bool
	Price                int
}

type CardPurchaseRecord struct {
	PurchaseId      int       `gorm:"primaryKey;autoIncrement"` // SERIAL PRIMARY KEY
	AdminAccount    string    `gorm:"size:255"`                 // VARCHAR(255)
	CardId          int       // card_lists(card_id)
	UserId          int       // INT
	Money           int       // INT
	InviteTeacherId int       // INT
	StartDate       time.Time // DATE
	EndDate         time.Time // DATE
	Days            int       // INT
	Times           int       // INT
}

type CardForbidList struct {
	ID       int `gorm:"primaryKey;autoIncrement"` // SERIAL PRIMARY KEY
	CardId   int // card_lists(card_id)
	CourseId int // courses(course_id)
}

type CardSupportList struct {
	ID       int `gorm:"primaryKey;autoIncrement"` // SERIAL PRIMARY KEY
	CardId   int // card_lists(card_id)
	CourseId int // courses(course_id)
}

type InputCardInfo struct {
	AdminAccount     string `json:"admin_account"`
	CardName         string `json:"card_name"`
	CardIntroduction string `json:"card_introduction"`
	IsSupportGroup   bool   `json:"is_support_group"`
	IsSupportTeam    bool   `json:"is_support_team"`
	IsSupportVIP     bool   `json:"is_support_vip"`
	IsLimitDays      bool   `json:"is_limit_days"`
	IsLimitTimes     bool   `json:"is_limit_times"`
	IsForbidSpecial  bool   `json:"is_forbid_special"`
	IsSupportSpecial bool   `json:"is_support_special"`
	ForbidCourseId   []int  `json:"forbid_course_id"`
	SupportCourseId  []int  `json:"support_course_id"`
	Price            int    `json:"price"`
}
