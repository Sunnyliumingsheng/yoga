package service

import (
	"fmt"
	"strconv"

	"api/db"
	"api/loger"
	"api/nets"
	"api/util"
	"api/weixin"
)

// 如果没有收到token或者session的时候调用,可能注册新用户,如果成功则返回一个db.Authentication,失败只会返回错误信息
func RegisterUser(code string) (message nets.Message) {
	openid, err := weixin.GetOpenId(code)
	if err != nil {
		return nets.Message{IsSuccess: false, Info: "获取openid失败", Result: err.Error()}
	}
	isExist, userId, level, err := db.IsThisOpenIdExistedAndGetLevel(openid)
	if err != nil {
		loger.Loger.Println("!!!!!!!!!!!!严重错误, 在检查用户是否存在的时候遇到了除了openid不存在以外的错误", err.Error(), "openid:", openid)
		return nets.Message{IsSuccess: false, Info: "获取用户信息失败", Result: err.Error()}
	}
	if isExist {
		// 给出新的token和session给到用户
		tokenChan := make(chan string)
		go util.AsyncGenerateToken(string(userId), tokenChan)
		go db.AddSession(fmt.Sprint(userId), level)
		token := <-tokenChan
		return nets.Message{IsSuccess: false, Info: "获取token成功", Result: token}
	} else {
		userId, err = db.InsertUserAndGetUserId(openid)
		if err != nil {
			loger.Loger.Println("!!!!!!!!!!!!严重错误, 检查了是否存在但仍然插入新用户失败", err.Error(), "openid:", openid)
			return nets.Message{IsSuccess: false, Info: "插入用户信息失败", Result: err.Error()}
		}
		tokenChan := make(chan string)
		go util.AsyncGenerateToken(string(userId), tokenChan)
		go db.AddSession(fmt.Sprint(userId), level)
		token := <-tokenChan
		return nets.Message{
			IsSuccess: true,
			Info:      "新用户注册成功",
			Result:    db.Authentication{Token: token, Session: fmt.Sprint(userId)},
		}
	}
}

// 如果客户端能提供session和token
func SessionAndTokenAuthentication(session string, token string) (message nets.Message) {
	userId, err := strconv.Atoi(session)
	if err != nil {
		return nets.Message{IsSuccess: false, Info: "session转换成int64失败", Result: err.Error()}
	}
	var isOnlineChan chan bool
	var levelChan chan int
	var isValidChan chan bool
	go db.AsyncAuthSession(session, isOnlineChan, levelChan)
	go util.AsyncParseToken(token, isValidChan)
	//一般来说redis的速度远快于token计算
	level := <-levelChan
	isOnline := <-isOnlineChan
	if isOnline {
		db.AddSession(session, level)
		return nets.Message{IsSuccess: true, Info: "session和验证通过,时间刷新", Result: nil}
	}
	isValid := <-isValidChan
	if isValid {
		level, err = db.IntUserIdSelectUserLevel(userId)
		if err != nil {
			loger.Loger.Println("!!!!!!!!!!!!严重错误, 在检查用户等级的时候遇到了err", err.Error(), "userId:", userId)
			return nets.Message{IsSuccess: false, Info: "获取用户等级失败", Result: err.Error()}
		}
		db.AddSession(session, level)
		return nets.Message{IsSuccess: true, Info: "token验证通过", Result: struct{}{}}
	}
	return nets.Message{IsSuccess: false, Info: "session和token不匹配", Result: nil}

}
