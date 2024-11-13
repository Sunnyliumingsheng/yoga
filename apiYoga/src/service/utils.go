package service

import "api/db"

func appendAdminInfo(admin db.Admin, user db.User) (info AdminInfo) {
	info.Account = admin.Account
	info.Name = user.Name
	info.AdminID = admin.AdminID
	info.AvaURL = user.AvaURL
	info.IsStudent = user.IsStudent
	info.IsTeacher = user.IsTeacher
	info.IsAdmin = true
	info.Nickname = user.Nickname
	info.Openid = user.Openid
	info.Password = admin.Password
	info.Signature = user.Signature
	info.Gender = user.Gender
	return info
}
func appendTeacherInfo(teacher db.Teacher, user db.User) (info TeacherInfo) {
	info.UserID = user.UserID
	info.Openid = user.Openid
	info.Nickname = user.Nickname
	info.Name = user.Name
	info.Gender = user.Gender
	info.Signature = user.Signature
	info.IsStudent = user.IsStudent
	info.IsTeacher = true
	info.IsAdmin = user.IsAdmin
	info.AvaURL = user.AvaURL
	info.TeacherID = teacher.TeacherID
	info.Account = teacher.Account
	info.Password = teacher.Password
	info.InviteCode = teacher.InviteCode
	info.Introduction = teacher.Introduction
	info.IntroductionURL = teacher.IntroductionURL
	return info
}
