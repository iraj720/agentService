package logic

import (
	"agents/configs"
	"fmt"
	"os"
	"time"
)

type AgentService struct {
	Cfg    configs.Config
	agents []Agent
}

func NewAgentService(cfg configs.Config) AgentService {
	agents := make([]Agent, 0)
	f, err := os.Create(cfg.AgentLogPath)
	if err != nil {
		panic(err)
	}
	for i, val := range cfg.AgentLocations {
		agents = append(agents, NewAgent(fmt.Sprintf("%d", i), val, f))
	}
	as := AgentService{Cfg: cfg, agents: agents}
	return as
}

func (s *AgentService) NewRequest(p []float32) {
	ba, _ := s.BestAgent(p)
	ba.IsBusy = true
	for {
		select {
		case ba.agentChan <- p:
			return
		case <-time.After(100 * time.Millisecond):
		}
	}
}

func (s *AgentService) StartRecieving() {
	for i := 0; i < s.Cfg.AgentsNumber; i++ {
		go func(a *Agent) {
			// point reader
			go func() {
				ticker := time.NewTicker(time.Second * 1)
				for {
					select {
					case <-ticker.C:
						a.lock.Lock()
						a.log(fmt.Sprintf("Im at point(x, y) : (%f, %f)", a.point[0], a.point[1]))
						a.lock.Unlock()
					}
				}
			}()

			// point writer
			for {
				select {
				case p := <-a.agentChan:
					a.log(fmt.Sprintf("Im Started going to point(x, y) : (%f, %f)", p[0], p[1]))
					a.lock.Lock()
					a.dist = p
					a.SetMoves()
					a.lock.Unlock()
					tc := time.NewTicker(1 * time.Second)
				l2:
					for {
						select {
						case <-time.After(time.Duration(time.Millisecond * time.Duration(a.findDistance()*1000))):
							a.lock.Lock()
							a.point = a.dist
							a.lock.Unlock()
							a.log(fmt.Sprintf("Im Ended my task at point(x, y) : (%f, %f) and distance is :%f ", a.point[0], a.point[1], a.findDistance()))
							break l2
						case <-tc.C:
							a.lock.Lock()
							a.UpdatePoint()
							a.lock.Unlock()
						}
					}
				}
			}
		}(&s.agents[i])
	}
}

func (s *AgentService) BestAgent(p []float32) (Agent, float32) {
	bestAgent := s.agents[0]
	var distance float32 = 0
	for i, _ := range s.agents {
		if !s.agents[i].IsBusy {
			Newdistance := s.agents[i].findDistance2(p)
			if Newdistance < distance {
				distance = Newdistance
				bestAgent = s.agents[i]
			}
		}
	}
	return bestAgent, distance
}
