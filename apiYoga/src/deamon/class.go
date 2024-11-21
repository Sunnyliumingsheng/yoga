package deamon

import (
	"api/db"
	"api/loger"
	"encoding/json"
	"os"
	"time"
)

type mymap map[db.ClassActived][]db.UserResumeInfo

var activedClass [4]mymap
var nowWeekDay int
var pmap int

type Savable struct {
	ClassInfo  map[int]db.ClassActived
	ResumeInfo map[int][]db.UserResumeInfo
	Num        int
}

func FlashActivedClass() {
	// 保存为记录 等下再写
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
}

// 将一个mymap转化为可保存的格式
func (m mymap) untie() (savable Savable) {
	savable.Num = 0
	savable.ClassInfo = make(map[int]db.ClassActived)
	savable.ResumeInfo = make(map[int][]db.UserResumeInfo)
	for class, resume := range m {
		savable.ClassInfo[savable.Num] = class
		savable.ResumeInfo[savable.Num] = resume
		savable.Num++
	}
	return savable
}
func (m mymap) tie(savable Savable) {
	m = nil
	m = make(mymap)
	for i := 0; i < savable.Num; i++ {
		m[savable.ClassInfo[i]] = savable.ResumeInfo[i]
	}
}

// 保存为json文件格式
func (m mymap) storage(savable Savable) {
	jsonData, err := json.Marshal(savable)
	if err != nil {
		loger.Loger.Println("error:不能解析json", err.Error())
		return
	}
	file, err := os.Create("resume.json")
	if err != nil {
		loger.Loger.Println("error:无法创建文件", err.Error())
		return
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		loger.Loger.Println("error:写入文件失败", err.Error())
		return
	}
	loger.Loger.Println("success: 已保存为json格式")
}
func (m mymap) recover() (savable Savable) {
	file, err := os.Open("resume.json")
	if err != nil {
		loger.Loger.Println("error: 无法打开json文件", err.Error())
		return
	}
	defer file.Close()
	jsonData := make([]byte, 1024)
	n, err := file.Read(jsonData)
	if err != nil {
		loger.Loger.Println("error: 读取json数据失败", err.Error())
		return
	}
	err = json.Unmarshal(jsonData[:n], &savable)
	if err != nil {
		loger.Loger.Println("error: 解析json数据失败", err.Error())
		return
	}
	return savable
}
