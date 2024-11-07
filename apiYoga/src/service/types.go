package service

// 在实际代码任务中，需要一些组合好的结构体，这样返回的可读性会更好
type AdminInfo struct {
	UserID    int
	Openid    string
	Nickname  string
	Name      string
	Gender    bool
	Signature string
	IsStudent bool
	IsTeacher bool
	AvaURL    string
	AdminID   int
	Account   string
	Password  string
}

type TeacherInfo struct {
	UserID          int
	Openid          string
	Nickname        string
	Name            string
	Gender          bool
	Signature       string
	IsStudent       bool
	IsAdmin         bool
	AvaURL          string
	TeacherID       int
	Account         string
	Password        string
	InviteCode      string
	Introduction    string
	IntroductionURL string
}
