package taskrunner

import (
	"errors"
	"testing"
)

func TestRunner_StartDispatch(t *testing.T) {
	d := func(d dataChan) error {
		for i := 0; i < 30; i++ {
			d <- i
			t.Logf("Dispatcher send: %v", i)
		}
		return nil
	}

	e := func(d dataChan) error {
		for {
			select {
			case data := <-d:
				t.Logf("Executor received %v", data)
			default:
				t.Logf("no data")
				return errors.New("execute is complete")
			}
		}
	}
	runner := NewRunner(30, false, d, e)
	runner.startAll()
}
