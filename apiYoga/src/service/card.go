package service

import (
	"api/db"
	"api/deamon"
	"api/loger"
	"errors"
	"time"
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
func (m *Message) SelectAllCardBasicInfo() {
	cardinfo, err := db.SelectAllCardBasicInfo()
	if err != nil {
		m.HaveError = true
		m.Info = err.Error()
		m.IsSuccess = false
		return
	}
	m.HaveError = false
	m.Info = "查询成功"
	m.Result = cardinfo
	m.IsSuccess = true
}
func (m *Message) BuyCard(username string, cardId, userId int, money int, endDate time.Time, times int, invitedTeacherId int) {
	err := db.InsertNewPurchaseCard(username, cardId, userId, money, endDate, times, invitedTeacherId)
	if err != nil {
		m.HaveError = true
		m.Info = err.Error()
		m.IsSuccess = false
		return
	}
	m.HaveError = false
	m.Info = "购卡成功"
	m.IsSuccess = true
}
func (m *Message) SelectPurchaseRecord(userId int) {
	purchaseRecord, err := db.SelectPurchaseRecord(userId)
	if err != nil {
		m.HaveError = true
		m.Info = err.Error()
		m.IsSuccess = false
		return
	}
	m.HaveError = false
	m.Info = "查询成功"
	m.Result = purchaseRecord
	m.IsSuccess = true
}
func (m *Message) DeletePurchaseRecord(purchaseId int) {
	err := db.DeletePurchaseRecord(purchaseId)
	if err != nil {
		m.HaveError = true
		m.Info = err.Error()
		m.IsSuccess = false
		return
	}
	loger.Loger.Println("money!!! 尝试删除一个购买记录", purchaseId)
	m.HaveError = false
	m.Info = "删除成功"
	m.IsSuccess = true
}
func CanStudentReserveThisCourse(userId int, courseId int) (isOk bool, err error) {
	basicCardInfo, isExist := deamon.UserCard[userId]
	if !isExist {
		return false, errors.New("请先购买会员卡")
	}
	courseType, err := db.SelectCourseTypeByCourseId(courseId)
	if err != nil {
		return false, err
	}
	if !basicCardInfo.IsForbidSpecial { //如果没有特殊禁止
		switch courseType {
		case 0:
			if basicCardInfo.IsSupportVIP {
				return true, nil
			}
		case 1:
			if basicCardInfo.IsSupportTeam {
				return true, nil
			}
		case 2:
			if basicCardInfo.IsSupportGroup {
				return true, nil
			}
		default:
			loger.Loger.Println("并没有捕获到错误,但是课程类型异常")
			return false, errors.New("逻辑错误，请联系管理员")
		}
		//如果简单的判断不行那就去看看特殊支持
		return db.IsCourseIdInCardSpecialSupport(basicCardInfo.CardId, courseId)
	} else { // 如果有特殊禁止
		// 有特殊禁止一律不给通过
		isForbid, err := db.IsCourseIdInSpecialForbid(basicCardInfo.CardId, courseId)
		if err != nil {
			return false, err
		}
		if isForbid {
			return false, nil
		} else { //没有被特殊禁止掉
			//正常检索一次
			courseType, err := db.SelectCourseTypeByCourseId(courseId)
			if err != nil {
				return false, err
			}
			switch courseType {
			case 0:
				if basicCardInfo.IsSupportVIP {
					return true, nil
				}
			case 1:
				if basicCardInfo.IsSupportTeam {
					return true, nil
				}
			case 2:
				if basicCardInfo.IsSupportGroup {
					return true, nil
				}
			default:
				loger.Loger.Println("并没有捕获到错误,但是课程类型异常2")
				return false, errors.New("逻辑错误，请联系管理员")
			}
			return false, nil

		}
	}
}
