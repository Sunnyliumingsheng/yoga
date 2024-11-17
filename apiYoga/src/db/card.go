package db

import "api/loger"

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
		for supportInfo := range input.SupportCourseId {
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
