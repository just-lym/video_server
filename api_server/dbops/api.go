package dbops

import (
	"database/sql"
	"github.com/just-lym/video_server/api_server/defs"
	"github.com/just-lym/video_server/api_server/utils"
	"log"
	"time"
)

func AddCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("insert into video_info (id, author_id, name, disply_ctime) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()
	return &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}, nil
}

func AddNewComments(vid string, aid string, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIn, err := dbConn.Prepare("insert into comments (id, video_id, content, author_id) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIn.Exec(id, vid, content, aid)
	if err != nil {
		return err
	}
	defer stmtIn.Close()
	return nil
}

func ListComments(videoId string, from int, to int) ([]*defs.Comments, error) {
	stmtOut, err := dbConn.Prepare(`select commits.id, users.login_name,comments.content from comments inner join 
    users on comments.author_id = users.id where comment.video_id=? and comments.time > FROM_UNIXTIME(?) and
                                                 comments.time <= FROM_UNIXTIME(?)`)
	if err != nil {
		return nil, err
	}
	var res []*defs.Comments
	rows, err := stmtOut.Query(videoId, from, to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, name, content string
		err := rows.Scan(&id, &name, &name, &content)
		if err != nil {
			return nil, err
		}
		c := &defs.Comments{Id: id, VideoId: videoId, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
