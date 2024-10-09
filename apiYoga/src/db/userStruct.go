package db

type User struct {
	UserID    uint    `gorm:"primaryKey"`
	OpenID    string  `gorm:"not null;size:255"`
	Nickname  string  `gorm:"not null;size:255"`
	Name      *string `gorm:"size:255"`
	Gender    *bool
	Signature *string `gorm:"size:64"`
	IsStudent bool
	IsTeacher bool
	IsAdmin   bool
	AvaURL    *string `gorm:"size:255"`
}

type Admin struct {
	AdminID  uint `gorm:"primaryKey"`
	UserID   uint
	Account  string `gorm:"size:255"`
	Password string `gorm:"size:255"`
}

type Teacher struct {
	TeacherID       uint    `gorm:"primaryKey"`
	Account         string  `gorm:"size:255"`
	Password        string  `gorm:"size:255"`
	InviteCode      string  `gorm:"size:255"`
	Introduction    string  `gorm:"type:text"`
	IntroductionURL *string `gorm:"size:255"`
}

type Student struct {
	StudentID uint `gorm:"primaryKey"`
	UserID    uint
}
