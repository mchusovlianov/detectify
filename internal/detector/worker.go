package detector

import (
	"log"
	"sync"
)

func (s *Service) startWorkerPool(num int, wg *sync.WaitGroup) {
	wg.Add(num)
	for i := 0; i < num; i += 1 {
		go s.worker(wg)
	}
}

func (s *Service) worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := <-s.taskQueue
		if !ok {
			return
		}

		scanInfo := ScanInfo{
			UUID:   task.UUID,
			Result: make(map[string]bool),
		}

		for _, host := range task.HostList {
			for _, probe := range s.probes {
				ok := probe.Run(host)
				scanInfo.Result[host] = ok
			}
		}

		err := s.output.Set(scanInfo.UUID, scanInfo)
		if err != nil {
			log.Printf("Error when saving %s - %s", scanInfo.UUID, err.Error())
		}
	}
}
