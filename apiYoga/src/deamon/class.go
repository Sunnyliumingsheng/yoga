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

func SelectResumeInfo(classId int)

func SelectClass() (resumeable [4][]db.ClassActived) {
	for i := 0; i < 4; i++ {
		for classActived, _ := range activedClass[i] {
			resumeable[i] = append(resumeable[i], classActived)
		}
	}
	return resumeable
}
func Resume(userId, classId int) (isOk bool, err error) {
	resumeWeekDay, err := db.SelectWeekdayByClassId(classId)
	if err != nil {
		return false, err
	}
	p := ((((resumeWeekDay - nowWeekDay) % 7) + pmap) % 4)
	if activedClass[p][db.ClassActived{ClassId: classId}] == nil {
		return false, nil
	}
	activedClass[p][db.ClassActived{ClassId: classId}] = append(activedClass[p][db.ClassActived{ClassId: classId}], db.UserResumeInfo{
		UserId: userId,
		Status: 0,
	})
	return true, nil
}

// 这里是今天添加今天的课这种情况使用的
func InsertNewActivedClass(class db.ClassList) {
	var classActived db.ClassActived
	classActived.ClassId = class.ClassId
	classActived.Index = class.Index
	classActived.ResumeNum = 0
	classActived.TeacherId = class.TeacherId
	classActived.Max = class.Max
	// this algorithm is checked shortly ,may mistake sometimes
	p := ((((class.DayOfWeek - nowWeekDay) % 7) + pmap) % 4)
	activedClass[p][classActived] = make([]db.UserResumeInfo, class.Max)
}

// when program start to do,init the data structure
func InitActivedClass() {
	weekday := time.Now().Weekday()
	nowWeekDay = int(weekday)
	for i := 0; i <= 3; i++ {
		classList, err := db.SelectActivedClassThisWeekday((nowWeekDay + i) % 7)
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
			activedClass[i] = make(mymap)
			activedClass[i][classActived] = make([]db.UserResumeInfo, 0, classActived.Max)
		}
	}
	pmap = 0
}

// renew the data structure,every night todo
func RenewActivedClass() {
	nowWeekDay = ((nowWeekDay + 1) % 7)
	activedClass[pmap] = nil
	newActivedClassList, err := db.SelectActivedClassThisWeekday(nowWeekDay)
	if err != nil {
		loger.Loger.Println("error:", err.Error())
		return
	}
	for _, class := range newActivedClassList {
		var classActived db.ClassActived
		classActived.ClassId = class.ClassId
		classActived.Index = class.Index
		classActived.ResumeNum = 0
		classActived.TeacherId = class.TeacherId
		classActived.Max = class.Max
		activedClass[pmap][classActived] = make([]db.UserResumeInfo, 0, classActived.Max)
	}
	pmap = (pmap + 1) % 4
}

type Savable struct {
	ClassInfo  map[int]db.ClassActived
	ResumeInfo map[int][]db.UserResumeInfo
	Num        int
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
func storage(savables [4]Savable) {
	jsonData, err := json.Marshal(savables)
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
func recover() (savables [4]Savable) {
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
	err = json.Unmarshal(jsonData[:n], &savables)
	if err != nil {
		loger.Loger.Println("error: 解析json数据失败", err.Error())
		return
	}
	return savables
}

// if you want to pause this program,and you are sure you will reboot soon,you can storage,and soon getStorage
func StorageClass() {
	var savables [4]Savable
	for i := 0; i <= 3; i++ {
		savables[i] = activedClass[i].untie()
	}
	storage(savables)
}
func GetStorageClass() {
	var savables [4]Savable
	savables = recover()
	for i := 0; i <= 3; i++ {
		activedClass[i] = nil
		activedClass[i].tie(savables[i])
	}
}
