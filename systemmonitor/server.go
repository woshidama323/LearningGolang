package systemmonitor

import (
	"log"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	hardwaremonitoring "github.com/woshidama323/LearningGolang/grpcserver"
)

type Server struct {
	hardwaremonitoring.UnimplementedHardwareMonitorServer
}

func (s *Server) Monitor(req *hardwaremonitoring.EmptyRequest, stream hardwaremonitoring.HardwareMonitor_MonitorServer) error {
	// Start a ticker that executes each 2 seconds
	timer := time.NewTicker(2 * time.Second)

	for {
		select {
		// Exit on stream context done
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			// Grab stats and output
			hwStats, err := s.GetStats()
			if err != nil {
				log.Println(err.Error())
			} else {

			}
			// Send the Hardware stats on the stream
			err = stream.Send(hwStats)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

func (s *Server) GetStats() (*hardwaremonitoring.HardwareStats, error) {
	// Extarcyt Memory statas
	mem, err := memory.Get()
	if err != nil {
		return nil, err
	}
	// Extract CPU stats
	cpu, err := cpu.Get()
	if err != nil {
		return nil, err
	}
	// Create our response object
	hwStats := &hardwaremonitoring.HardwareStats{
		Cpu:        int32(cpu.Total),
		MemoryFree: int32(mem.Free),
		MemoryUsed: int32(mem.Used),
	}

	return hwStats, nil
}
