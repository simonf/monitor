package main

import (
        "fmt"
	"os"
	"simonf.net/monitor_agent"
	"simonf.net/monitor_listener"
	"time"
)

func main() {
     	if len(os.Args) != 2 {
     	   usage();
	} else {
	   if os.Args[1] == "-s" {
		runAsServer()
	   } else if os.Args[1] == "-a" {
	   	runAsAgent()
	   } else {
	     	monitor_agent.RunOnce()
	   }
	}
}

func usage() {
     fmt.Println("Usage: monitor -[a|s|o]")
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
