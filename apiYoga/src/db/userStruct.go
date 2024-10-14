package db

type User struct {
	UserID    int    `gorm:"primaryKey"`
	Openid    string `gorm:"not null;size:255"`
	Nickname  string `gorm:"not null;size:255"`
	Name      string `gorm:"size:255"`
	Gender    bool
	Signature string `gorm:"size:64"`
	IsStudent bool
	IsTeacher bool
	IsAdmin   bool
	AvaURL    string `gorm:"size:255"`
}

type Admin struct {
	AdminID  int `gorm:"primaryKey"`
	UserID   int
	Account  string `gorm:"size:255"`
	Password string `gorm:"size:255"`
}

type Teacher struct {
	UserID          int    `gorm:"size:255"`
	TeacherID       int    `gorm:"primaryKey"`
	Account         string `gorm:"size:255"`
	Password        string `gorm:"size:255"`
	InviteCode      string `gorm:"size:255"`
	Introduction    string `gorm:"type:text"`
	IntroductionURL string `gorm:"size:255"`
}

type Student struct {
	StudentID int `gorm:"primaryKey"`
	UserID    int
}

type Authentication struct {
	Session string `json:"session"`
	Token   string `json:"token"`
}
