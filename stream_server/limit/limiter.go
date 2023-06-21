package limit

import "github.com/just-lym/video_server/stream_server/customlog"

type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		customlog.Errorln("Reached the rate limit")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	<-cl.bucket
	customlog.Infoln("New connection coming")
}
