package nets

import "github.com/gin-gonic/gin"

func basicalApiEngine(r *gin.Engine) {
	// 用户端使用的有关用户的
	r.POST("/api/login", userLoginWithSessionAndToken)
	r.POST("/api/register", userLoginWithCode)
	r.POST("/api/rename", userRename)
	// 开发者和管理员使用的有关用户的
	r.POST("/api/root/register/admin", sudoRegisterAdmin)
	r.POST("/api/root/login", sudoLogin)
	r.POST("/api/root/select/user/by/name", selectUserInfoByName)
	r.POST("/api/root/insert/user/by/name", insertNewUser)
	r.POST("/api/root/drop/user/by/userId", dropUserByUserId)
	r.POST("/api/root/update/user/level/by/name", updateUserLevel)
	r.POST("/api/root/select/user/tail", selectUserTail)
	r.POST("/api/root/insert/teacher", insertTeacherAccountAndPassword)
	r.POST("/api/root/insert/admin", insertAdminAccountAndPassword)
	r.POST("/api/root/select/admin/by/name", selectAdminInfoByName)
	r.POST("/api/root/select/teacehr/by/name", selectTeacherInfo)
	// 管理员和老师使用的有关用户的
	r.POST("/api/admin/login", adminAndTeacherLogin)
	// 管理员和老师使用的有关课程的
	r.POST("/api/admin/drop/course/by/name", dropCourseByName)
	r.POST("/api/admin/insert/course", insertNewCourse)

}
