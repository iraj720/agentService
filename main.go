package main

import (
	"agents/configs"
	"agents/internal/logic"
	"bufio"
	"io"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	cfg := configs.LoadConfigs()
	agentService := logic.NewAgentService(cfg)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM)
	TurnOnTherminal(agentService)
}

func TurnOnTherminal(as logic.AgentService) {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()
	go as.StartRecieving()

	for {
		customersRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var customersRow []float32
		for _, customersRowItem := range customersRowTemp {
			customersItemTemp, err := strconv.ParseFloat(customersRowItem, 10)
			checkError(err)
			customersItem := customersItemTemp
			customersRow = append(customersRow, float32(customersItem))
		}

		if len(customersRow) != 2 {
			panic("Bad input")
		}
		go as.NewRequest(customersRow)
	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
