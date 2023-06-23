package dbops

import "github.com/just-lym/video_server/scheduler/customlog"

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("select video_id from video_del_record limit ?")
	var ids []string
	if err != nil {
		customlog.Errorln("prepare find video deletion record error %v", err)
		return ids, err
	}

	query, err := stmtOut.Query(count)
	if err != nil {
		customlog.Errorln("query video deletion record error %s", err.Error())
		return ids, err
	}
	for query.Next() {
		var id string
		err := query.Scan(&id)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmtOut, err := dbConn.Prepare("delete from video_del_record where id = ?")
	if err != nil {
		return err
	}
	_, err = stmtOut.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtOut.Close()
	return nil
}
