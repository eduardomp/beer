package services

import (
	"beer/core/model"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/process"
)

func GetProcessInfo() []model.Process {
	c := make(chan model.Process)
	processes, _ := process.Processes()
	result := make([]model.Process, 0)

	for i := 0; i < len(processes); i++ {
		p := processes[i]
		go func() {
			status, _ := p.Status()
			name, _ := p.Name()
			cmd, _ := p.Cmdline()
			connections, _ := p.Connections()

			if len(connections) > 0 {

				c <- model.Process{
					Name:    name,
					Pid:     int(p.Pid),
					Cmdline: cmd,
					Status:  status,
				}

			}

		}()

	}

	for range c {
		select {
		case proc := <-c:
			result = append(result, proc)
		case <-time.After(time.Second * 10):
			fmt.Println("Timeout to collect process information!")
			close(c)
		}
	}

	return result
}
