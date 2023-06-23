package taskrunner

import (
	"errors"
	"github.com/just-lym/video_server/scheduler/customlog"
	"github.com/just-lym/video_server/scheduler/dbops"
	"os"
	"sync"
)

func VideoClearDispatcher(data dataChan) error {
	records, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		customlog.Errorln("Video clear dispatcher error %v", err)
		return err
	}
	if len(records) <= 0 {
		return errors.New("all tasks finished")
	}
	for _, record := range records {
		data <- record
	}
	return nil
}

func VideoClearExecutor(data dataChan) error {
	group := sync.WaitGroup{}
	errMap := sync.Map{}
	var err error
forLoop:
	for {
		select {
		case vid := <-data:
			group.Add(1)
			go func(waitGroup *sync.WaitGroup, vid string) {
				_, err := os.ReadFile("./video/" + vid)
				if err != nil && !os.IsNotExist(err) {
					customlog.Errorln("deleting video file error")
					errMap.Store(vid, err)
					return
				}
				err = dbops.DelVideoDeletionRecord(vid)
				if err != nil {
					customlog.Errorln("deleting record error")
					errMap.Store(vid, err)
				}
				group.Done()
			}(&group, vid.(string))
		default:
			break forLoop
		}
	}
	group.Wait()
	errMap.Range(func(key, value any) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
