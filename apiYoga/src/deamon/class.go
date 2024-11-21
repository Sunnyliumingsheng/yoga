package deamon

import (
	"api/db"
	"api/loger"
	"time"
)

// int 数组里面是关于已经预约的用户的信息
var activedClass map[db.ClassActived][]db.UserResumeInfo

func FlashActivedClass() {
	// 保存为记录 等下再写
	activedClass = nil
	// 找到所有已经激活的而且dayofweek相等于今天的
	weekday := time.Now().Weekday()
	intWeekday := int(weekday)
	classList, err := db.SelectActivedClassThisWeekday(intWeekday)
	if err != nil {
		loger.Loger.Println("error:", err.Error())
	}
	for _, class := range classList {
		var classActived db.ClassActived
		classActived.ClassId = class.ClassId
		classActived.Index = class.Index
		classActived.ResumeNum = 0
		classActived.TeacherId = class.TeacherId
		classActived.Max = class.Max
		// 这里使用make如果不设置0,就会自动填充0值，而max指定有利于性能
		activedClass[classActived] = make([]db.UserResumeInfo, 0, class.Max)
	}
}

// 这里是今天添加今天的课这种情况使用的
func AddActivedClass(class db.ClassList) {
	var classActived db.ClassActived
	classActived.ClassId = class.ClassId
	classActived.Index = class.Index
	classActived.ResumeNum = 0
	classActived.TeacherId = class.TeacherId
	classActived.Max = class.Max
	// 这里使用make如果不设置0,就会自动填充0值，而max指定有利于性能
	activedClass[classActived] = make([]db.UserResumeInfo, 0, class.Max)
}
