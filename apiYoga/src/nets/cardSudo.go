package nets

import (
	"api/service"
	"time"

	"github.com/gin-gonic/gin"
)

func buyCard(c *gin.Context) {
	type newPurchaseInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		CardId             int                `json:"card_id"`
		UserId             int                `json:"user_id"`
		Money              int                `json:"money"`
		InviteTeacherId    int                `json:"invite_teacher_id"`
		EndDate            time.Time          `json:"end_date"`
		Times              int                `json:"times"`
	}
	var getData newPurchaseInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"error": "验证失败"})
		return
	}
	var m service.Message
	m.BuyCard(getData.SudoAuthentication.SudoName, getData.CardId, getData.UserId, getData.Money, getData.EndDate, getData.Times, getData.InviteTeacherId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
func deletePurchaseRecordByPurchaseId(c *gin.Context) {
	type deleteInfo struct {
		SudoAuthentication SudoAuthentication `json:"sudoAuthentication"`
		PurchaseId         int                `json:"purchase_id"`
	}
	var getData deleteInfo
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !authenticateSudo(getData.SudoAuthentication) {
		c.JSON(400, gin.H{"error": "验证失败"})
		return
	}
	var m service.Message
	m.DeletePurchaseRecord(getData.PurchaseId)
	if m.HaveError {
		c.JSON(400, gin.H{"error": m.Info})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
