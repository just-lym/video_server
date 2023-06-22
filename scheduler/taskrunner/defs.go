package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READ_TO_EXECUTE   = "e"
	CLOSE             = "c"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(data dataChan) error
