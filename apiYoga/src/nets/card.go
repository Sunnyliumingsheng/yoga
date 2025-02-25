package nets

import (
	"api/db"
	"api/service"

	"github.com/gin-gonic/gin"
)

func insertNewCard(c *gin.Context) {
	type cardInfo struct {
		Token            string `json:"token"`
		CardName         string `json:"card_name"`
		CardIntroduction string `json:"card_introduction"`
		IsSupportGroup   bool   `json:"is_support_group"`
		IsSupportTeam    bool   `json:"is_support_team"`
		IsSupportVIP     bool   `json:"is_support_vip"`
		IsLimitDays      bool   `json:"is_limit_days"`
		IsLimitTimes     bool   `json:"is_limit_times"`
		IsForbidSpecial  bool   `json:"is_forbid_special"`
		IsSupportSpecial bool   `json:"is_support_special"`
		ForbidCourseId   []int  `json:"forbid_course_id"`
		SupportCourseId  []int  `json:"support_course_id"`
		Price            int    `json:"price"`
	}
	var getData cardInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, htTeacher, account := authenticationInElectron(getData.Token, c)
	if !isOk {
		c.JSON(400, gin.H{"error": "验证错误"})
	}
	if htTeacher == false {
		c.JSON(400, gin.H{"message": "只有管理员才能添加新卡"})
		return
	}
	inputCardInfo := db.InputCardInfo{
		AdminAccount:     account,
		CardName:         getData.CardName,
		CardIntroduction: getData.CardIntroduction,
		IsSupportGroup:   getData.IsSupportGroup,
		IsSupportTeam:    getData.IsSupportTeam,
		IsSupportVIP:     getData.IsSupportVIP,
		IsLimitDays:      getData.IsLimitDays,
		IsLimitTimes:     getData.IsLimitTimes,
		IsForbidSpecial:  getData.IsForbidSpecial,
		IsSupportSpecial: getData.IsSupportSpecial,
		ForbidCourseId:   getData.ForbidCourseId,
		SupportCourseId:  getData.SupportCourseId,
		Price:            getData.Price,
	}
	var m service.Message
	m.InsertNewCard(inputCardInfo)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}

// 删除会员卡，如果已经有人买了这个卡会失效的
func deleteNewCardByName(c *gin.Context) {
	type deleteCardInfo struct {
		Token    string `json:"token"`
		CardName string `json:"card_name"`
	}
	var getData deleteCardInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	isOk, htTeacher, account := authenticationInElectron(getData.Token, c)
	if !isOk {
		c.JSON(400, gin.H{"error": "验证错误"})
		return
	}
	if htTeacher == false {
		c.JSON(400, gin.H{"message": "只有管理员才能删除会员卡"})
		return
	}
	var m service.Message
	m.DeleteNewCardByName(getData.CardName, account)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// 这里做出了修改，注释可见，即为解除了查看的权限
func selectAllCardBasicInfo(c *gin.Context) {
	// type userInfo struct {
	// 	AuthenticationInfo AuthenticationInfo `json:"authenticationInfo"`
	// }
	// var getData userInfo
	// if err := c.ShouldBindJSON(&getData); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// _, err := authentication(getData.AuthenticationInfo, c)
	// if err != nil {
	// 	return
	// }
	var m service.Message
	m.SelectAllCardBasicInfo()
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	var cards []db.CardComplexInfo
	cards, ok := m.Result.([]db.CardComplexInfo)
	if !ok {
		c.JSON(400, gin.H{"error": "解析返回结果失败"})
	}
	c.JSON(200, gin.H{"cards": cards})
}

// 这个可以查看某人的购买记录，为了方便也不需要鉴定权限
func selectPurchaseRecordByUserId(c *gin.Context) {
	type selectInfo struct {
		UserId int `json:"user_id"`
	}
	var getData selectInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var m service.Message
	m.SelectPurchaseRecord(getData.UserId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	var purchaseRecords []db.CardPurchaseRecord
	purchaseRecords, ok := m.Result.([]db.CardPurchaseRecord)
	if !ok {
		c.JSON(400, gin.H{"error": "解析返回结果失败"})
	}
	c.JSON(200, gin.H{"purchase_records": purchaseRecords})
}
