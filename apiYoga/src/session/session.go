package session

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

type sessionInfo struct {
	userID    int
	level     int
	expiredAt time.Time
}

type sessionManager struct {
	sessions map[string]sessionInfo
	mu       sync.Mutex
	expiry   time.Duration
}

var m sessionManager

func init() {
	m.sessions = make(map[string]sessionInfo)
	m.mu = sync.Mutex{}
	m.expiry = 2 * time.Second
	go cleanExpiredSessionInfo(m.expiry)
}

func InsertSession(userId, level int) (sessionId string) {
	newSessionId := generateSessionID()
	now := time.Now()
	expiredAt := now.Add(m.expiry)
	m.mu.Lock()
	m.sessions[newSessionId] = sessionInfo{
		userID:    userId,
		level:     level,
		expiredAt: expiredAt,
	}
	m.mu.Unlock()
	return newSessionId
}
func CheckSession(sessionId string) (isOk bool, userId, level int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	sessionInfo, isExist := m.sessions[sessionId]
	now := time.Now()
	if !isExist {
		return false, -1, -1
	}
	if sessionInfo.expiredAt.Before(now) {
		//已经过期
		return false, -1, -1
	} else {
		//为了避免用户经常性的使用token,通过的话就刷新session的过期时间
		sessionInfo.expiredAt = now.Add(m.expiry)
		return true, sessionInfo.userID, sessionInfo.level
	}
}

// generateSessionID 生成一个随机的 session ID
func generateSessionID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic("无法生成随机 session ID")
	}
	return hex.EncodeToString(b)
}
func cleanExpiredSessionInfo(expiry time.Duration) {
	for range time.Tick(2 * expiry) {
		m.mu.Lock()
		now := time.Now()
		for sessionId, sessionInfo := range m.sessions {
			if sessionInfo.expiredAt.Before(now) {
				//过期
				delete(m.sessions, sessionId)
			}
		}
		m.mu.Unlock()
	}
}
