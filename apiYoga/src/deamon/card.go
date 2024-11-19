package deamon

import "api/db"

// 这是个有关会员卡的线程，用于将用户会员卡的信息存储在内存中，减少查询和处理的时间

// 这个是大多数面相用户使用时的会员卡信息，也就是说类似约课之类的操作，只会走这里过，而不是直接去数据库里找
var UserCard map[int]db.BasicCardInfo

// 调用一次这个接口，就会刷新一次上面的这个map
func FlashUserCard() {
	UserCard = nil
	UserCard = make(map[int]db.BasicCardInfo)
	err := db.SelectBasicCardInfo(UserCard)
	if err != nil {
		panic(err.Error())
	}
}

// 为减少性能消耗，新卡购买完全可以不需要使用太多的资源，只添加一个新的就行了
func FlashUserCardForNewCard(userId int, cardInfo db.BasicCardInfo) {
	UserCard[userId] = cardInfo
}
