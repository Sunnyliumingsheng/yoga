package deamon

import (
	"api/db"
)

// int 数组里面是关于已经预约的用户的信息
var activedClass map[db.ClassActived][]db.UserResumeInfo

func FlashActivedClass() {
	activedClass = nil

}
func AddActivedClass() {

}
