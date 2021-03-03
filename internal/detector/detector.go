package detector

import (
	"github.com/mchusovlianov/detectify/internal/probe/nginx"
	"github.com/mchusovlianov/detectify/internal/storage"
	"sync"
)

type Service struct {
	taskQueue chan Task
	probes    []Probe
	output    Storage
}

func NewService(taskSize int) *Service {
	s := &Service{
		taskQueue: make(chan Task, taskSize),
		probes: []Probe{
			&nginx.NginxProbe{},
		},
		output: storage.NewMemoryStorage(),
	}

	return s
}

func (s *Service) Run() error {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	s.startWorkerPool(4, wg)
	go startHttpServer(s, wg)

	wg.Wait()

	return nil
}
