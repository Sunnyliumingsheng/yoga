package nets

import "github.com/gin-gonic/gin"

func basicalApiEngine(r *gin.Engine) {
	// 照片相关的
	r.POST("/api/upload/picture", uploadPicture)
	r.POST("/api/delete/picture", deletePicture)
	// 用户端使用的有关用户的
	r.POST("/api/login", userLoginWithSessionAndToken)
	r.POST("/api/register", userLoginWithCode)
	r.POST("/api/rename", userRename)
	r.POST("/api/update/user/info", updateUserInfo)
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
	r.POST("/api/root/select/course", selectCourse)
	// 管理员和老师使用的有关用户的
	r.POST("/api/admin/login", adminAndTeacherLogin)
	r.POST("/api/admin/entrance", electronEntrance)
	r.POST("/api/admin/update/teacher/info", updateTeacherInfo)
	// 管理员和老师使用的有关课程的
	r.POST("/api/admin/delete/course/by/name", deleteCourseByName)
	r.POST("/api/admin/insert/course", insertNewCourse)
	r.POST("/api/admin/select/all/course", selectCourseByElectron)
	// 用户端使用的关于课程的
	r.POST("/api/select/all/course", selectCourse)
	// 任何人使用的有关课程的
	r.POST("/api/select/course", selectCourseInfo)
	// 管理员和老师使用的有关会员卡的
	r.POST("/api/admin/delete/card/by/name", deleteNewCardByName)
	r.POST("/api/admin/insert/card", insertNewCard)
	r.POST("/api/admin/select/all/card", selectAllCardBasicInfo)
	// 开发者和管理员使用的
	r.POST("/api/root/delete/purchase/record/by/purchaseId", deletePurchaseRecordByPurchaseId)
	// 所有人都可以使用的检索购买记录的
	r.POST("/api/select/purchase/record/by/userId", selectPurchaseRecordByUserId)
	// 管理员和老师使用的有关课程的
	r.POST("/api/admin/insert/class", insertNewClass)
	r.POST("/api/admin/delete/class", deleteClass)
	r.POST("/api/admin/active/class", activeClass)
	r.POST("/api/admin/select/all/class", selectAllClass)
	//
	r.POST("/api/select/all/actived/class", selectAllActivedClass)
	r.POST("/api/select/actived/class", selectActivedClass)
	r.POST("/api/select/class/by/classId", selectClassByClassId)
	//
	r.POST("/api/resume", resume)
	r.POST("/api/select/teaching/class", selectTeachingClass)
	r.POST("/api/cancel/resume", cancelResume)
	r.POST("/api/select/my/resume", selectMyResume)
	r.POST("/api/update/checkin/Status/by/userId", updateTeacherInfoCheckinStatusByUserId)
	r.POST("/api/select/record", selectRecord)
	//
	r.POST("/api/admin/select/black/list/by/userId", selectBlackListByUserId)
	r.POST("/api/admin/delete/all/black/list", deleteAllBlackList)
	r.POST("/api/admin/select/black/list", SelectBlackList)
	// thanks the nginx ,no one can access this api,without I testing
	r.POST("/test/storage/ram")
	r.POST("/test/get/storage/ram")
}
