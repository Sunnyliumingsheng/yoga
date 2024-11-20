package deamon

import "api/loger"

func StartAllDeamon() {
	FlashUserCard()
	loger.Loger.Println("finish all deamon")
}
