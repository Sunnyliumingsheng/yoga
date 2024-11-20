package nets

import (
	"api/db"
	"api/service"

	"github.com/gin-gonic/gin"
)

func insertNewClass(c *gin.Context) {
	type classInfo struct {
		ClassName     string `json:"class_name"`
		CourseId      int    `json:"course_id"`
		Auto          bool   `json:"auto"`
		DayOfWeek     int    `json:"day_of_week"`
		AlreadyActive bool   `json:"already_active"`
		Index         int    `json:"index"`
		Min           int    `json:"min"`
		Max           int    `json:"max"`
		TeacherId     int    `json:"teacher_id"`
		Token         string `json:"token"`
	}
	var getData classInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, isAdmin, _ := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	if !isAdmin {
		c.JSON(400, gin.H{"message": "只有管理员才能创建新班级"})
		return
	}
	var classList db.ClassList
	classList.ClassName = getData.ClassName
	classList.CourseId = getData.CourseId
	classList.Auto = getData.Auto
	classList.DayOfWeek = getData.DayOfWeek
	classList.AlreadyActive = getData.AlreadyActive
	classList.Index = getData.Index
	classList.Min = getData.Min
	classList.Max = getData.Max
	classList.TeacherId = getData.TeacherId
	var m service.Message
	m.InsertNewClass(classList)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}
func deleteClass(c *gin.Context) {
	type deleteInfo struct {
		ClassId int    `json:"class_id"`
		Token   string `json:"token"`
	}
	var getData deleteInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, isAdmin, _ := authenticationInElectron(getData.Token, c)
	if !isOk {
		return
	}
	if !isAdmin {
		c.JSON(400, gin.H{"message": "只有管理员才能删除课程"})
		return
	}
	var m service.Message
	m.DeleteClass(getData.ClassId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": m.Info})
}
