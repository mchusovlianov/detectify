package detector

import (
	"errors"
	"time"
)

type Task struct {
	UUID     string
	HostList []string
	ProbeIDs []string
}

func (s *Service) addToTaskQueue(task Task) error {
	timeout := time.NewTimer(time.Millisecond * 50)
	syncCh := make(chan struct{})
	go func() {
		s.taskQueue <- task
		syncCh <- struct{}{}
	}()

	select {
	case <-timeout.C:
		return errors.New("Timeout")
	case <-syncCh:
	}

	return nil
}
