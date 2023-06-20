package session

import (
	"github.com/just-lym/video_server/api_server/dbops"
	"github.com/just-lym/video_server/api_server/defs"
	"github.com/just-lym/video_server/api_server/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func LoadSessionFromDB() {
	session, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}
	session.Range(func(key, value any) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	uuid, _ := utils.NewUUID()
	ctime := time.Now().UnixNano() / 1000000
	ttl := ctime + 30*60*1000
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(uuid, ss)
	dbops.InsertSession(uuid, ttl, un)
	return uuid
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func IsSessionExpired(sid string) (string, bool) {
	value, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if value.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return value.(defs.SimpleSession).Username, false
	}
	return "", true
}
