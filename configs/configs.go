package configs

import "math/rand"

type Config struct {
	AgentsNumber   int
	AgentLocations [][]float32
	AgentLogPath   string
}

func LoadConfigs() Config {
	cfg := Config{AgentsNumber: 2}
	agentLocations := make([][]float32, 0)
	for i := 0; i < cfg.AgentsNumber; i++ {
		Xaxis := rand.Float32() * 10
		Yaxis := rand.Float32() * 10
		agentLocations = append(agentLocations, []float32{Xaxis, Yaxis})
	}
	cfg.AgentLocations = agentLocations
	cfg.AgentLogPath = "./agentsLog"
	return cfg
}
