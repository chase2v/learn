package session

import (
	"time"
	"sync"
	"github.com/avenssii/video_server/api/defs"
	"github.com/avenssii/video_server/api/dbops"
)

type SimpleSession struct {
	Username string // login anme
	TTL int64
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func deleteExpiredSession(sid string) {

}

func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(k, v interface{} bool{
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.newUUID()
	ct := time.Now().UnixNano()/100000
	ttl := ct + 30 * 60 * 1000 // Serverside session valid time: 30 min

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano()/1000000
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expired session
			return "", true
		}

		return ss.(*defs.SimpleSeesion).Username, false
	}

	return "", true
}

