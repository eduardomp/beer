package services

import (
	"beer/core/model"
	"sync"

	"github.com/shirou/gopsutil/process"
)

func GetProcessInfo() []model.Process {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	processes, _ := process.Processes()
	result := make([]model.Process, 0)

	for i := 0; i < len(processes); i++ {
		p := processes[i]
		wg.Add(1)
		go func() {
			defer wg.Done()

			status, _ := p.Status()
			name, _ := p.Name()
			cmd, _ := p.Cmdline()
			connections, _ := p.Connections()

			if len(connections) > 0 {
				mutex.Lock()
				result = append(result, model.Process{
					Name:    name,
					Pid:     int(p.Pid),
					Cmdline: cmd,
					Status:  status,
				})
				mutex.Unlock()
			}

		}()

	}

	wg.Wait()

	return result
}
