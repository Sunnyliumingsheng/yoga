package db

import (
	"api/loger"
	"time"
)

func InsertNewCard(input InputCardInfo) (err error) {
	tx := postdb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	cardInfo := CardList{
		CardName:         input.CardName,
		CardIntroduction: input.CardIntroduction,
		IsSupportGroup:   input.IsSupportGroup,
		IsSupportTeam:    input.IsSupportTeam,
		IsSupportVIP:     input.IsSupportVIP,
		IsLimitDays:      input.IsLimitDays,
		IsLimitTimes:     input.IsLimitTimes,
		IsForbidSpecial:  input.IsForbidSpecial,
		IsSupportSpecial: input.IsSupportSpecial,
		AdminAccount:     input.AdminAccount,
		Price:            input.Price,
	}
	if err := tx.Error; err != nil {
		return err
	}
	err = tx.Model(&CardList{}).Create(&cardInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	loger.Loger.Println("新插入了一个", cardInfo.CardName, cardInfo.CardId)
	if cardInfo.IsForbidSpecial {
		for forbidInfo := range input.ForbidCourseId {
			cardForbidInfo := CardForbidList{
				CourseId: forbidInfo,
				CardId:   cardInfo.CardId,
			}
			err = tx.Model(&CardForbidList{}).Create(cardForbidInfo).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	if cardInfo.IsSupportSpecial {
		for _, supportInfo := range input.SupportCourseId {
			cardSupportInfo := CardSupportList{
				CourseId: supportInfo,
				CardId:   cardInfo.CardId,
			}
			err = tx.Model(&CardSupportList{}).Create(cardSupportInfo).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
func DeleteNewCardByName(cardName string) (err error) {
	err = postdb.Where("card_name = ?", cardName).Delete(&CardList{}).Error
	return err
}
func SelectAllCardBasicInfo() (cardInfo []CardComplexInfo, err error) {
	// 检索功能是不用事务的
	var cardLists []CardList
	err = postdb.Model(&CardList{}).Find(&cardLists).Error
	if err != nil {
		return nil, err
	}
	for index, card := range cardLists {
		var forbidCourses []CourseBasic
		var supportCourse []CourseBasic
		err = postdb.Model(&CardForbidList{}).Where("card_id=?", card.CardId).Select("course_id,course_name").Find(&forbidCourses).Error
		if err != nil {
			return nil, err
		}
		cardInfo[index].ForbidCourseInfo = forbidCourses
		err = postdb.Model(&CardSupportList{}).Where("card_id=?", card.CardId).Select("course_id,course_name").Find(&supportCourse).Error
		if err != nil {
			return nil, err
		}
		cardInfo[index].SupportCourseInfo = supportCourse
		cardInfo[index].CardInfo = card
	}
	return cardInfo, nil
}
func InsertNewPurchaseCard(username string, cardId, userId int, money int, endDate time.Time, times int, invitedTeacherId int) (err error) {
	newPurchaseInfo := CardPurchaseRecord{
		AdminUsername:   username,
		CardId:          cardId,
		UserId:          userId,
		Money:           money,
		InviteTeacherId: invitedTeacherId,
		StartDate:       time.Now(),
		EndDate:         endDate,
		Times:           times,
	}
	err = postdb.Model(&CardPurchaseRecord{}).Create(&newPurchaseInfo).Error
	return err
}
func SelectBasicCardInfo(UserCard map[int]BasicCardInfo) (err error) {
	currency := time.Now()
	var userIds []int
	err = postdb.Model(&CardPurchaseRecord{}).Where("end_date>=?", currency).Pluck("user_id", &userIds).Error
	if err != nil {
		loger.Loger.Println("money!!!:在检索用户购买会员卡的列表的时候严重问题,error:", err.Error())
		return err
	}
	for _, userId := range userIds {
		var cardInfo CardList
		err = postdb.Model(&CardList{}).Where("user_id=?", userId).First(&cardInfo).Error
		if err != nil {
			loger.Loger.Println("money!!!:", "在检索用户购买列表的时候检索失败userid:", userId, "error:", err.Error())
			return err
		}
		UserCard[userId] = BasicCardInfo{
			CardId:           cardInfo.CardId,
			IsSupportGroup:   cardInfo.IsSupportGroup,
			IsSupportTeam:    cardInfo.IsSupportTeam,
			IsSupportVIP:     cardInfo.IsSupportVIP,
			IsLimitDays:      cardInfo.IsLimitDays,
			IsLimitTimes:     cardInfo.IsLimitTimes,
			IsForbidSpecial:  cardInfo.IsForbidSpecial,
			IsSupportSpecial: cardInfo.IsSupportSpecial,
		}
	}
	return nil
}
