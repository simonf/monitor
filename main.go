package main

import (
	"os"
	"simonf.net/monitor_agent"
	"simonf.net/monitor_listener"
	"time"
)

func main() {
	if os.Args[1] == "-s" {
		runAsServer()
	} else {
		runAsAgent()
	}
}

func runAsServer() {
	go monitor_listener.ListenForClients()
	go monitor_listener.PeriodicallyAdvertise()
	monitor_listener.StartServer(8001)
}

func runAsAgent() {
	go monitor_agent.Start()
	for {
		// wait for something to happen
		time.Sleep(1 * time.Minute)
	}
}
