package dbops

import (
	"database/sql"
	"github.com/just-lym/video_server/api_server/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("insert into sessions (session_id, TTL, login_name) values (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL, login_name from sessions where session_id = ?;")
	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, err
}

func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlStr string
		var loginName string
		if err := rows.Scan(&id, &ttlStr, &loginName); err != nil {
			log.Printf("retrieve session error:%s", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlStr, 10, 64); err == nil {
			ss := &defs.SimpleSession{Username: loginName, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id:%s, ttl:%d")
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmt, err := dbConn.Prepare("delete from sessions where session_id =?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil
}
