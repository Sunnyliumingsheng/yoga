package service

import (
	"api/db"
	"api/loger"
)

func (m *Message) InsertNewCard(input db.InputCardInfo) {
	err := db.InsertNewCard(input)
	if err != nil {
		m.HaveError = true
		m.Info = err.Error()
		m.IsSuccess = false
		return
	}
	m.HaveError = false
	m.Info = "插入成功"
	m.IsSuccess = true
}
func (m *Message) DeleteNewCardByName(cardName string, account string) {
	loger.Loger.Println(account, "account的管理员尝试删除一个会员卡类型", cardName)
	err := db.DeleteNewCardByName(cardName)
	if err != nil {
		m.HaveError = true
		m.Info = err.Error()
		m.IsSuccess = false
		return
	}
	m.HaveError = false
	m.Info = "删除成功"
	m.IsSuccess = true
	return
}
