package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/fly-apps/mongo-flex/internal/supervisor"
)

func main() {
	log.SetFlags(0)

	svisor := supervisor.New("flymongo", 5*time.Minute)

	svisor.AddProcess("mongod", "mongod --config /etc/mongod.conf")
	svisor.AddProcess("admin", "/usr/local/bin/start_admin_server",
		supervisor.WithRestart(0, 5*time.Second),
	)

	svisor.StopOnSignal(syscall.SIGINT, syscall.SIGTERM)

	if err := svisor.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func panicHandler(err error) {
	debug := os.Getenv("DEBUG")
	if debug != "" {
		fmt.Println(err.Error())
		fmt.Println("Entering debug mode... (Timeout: 10 minutes")
		time.Sleep(time.Minute * 10)
	}

	panic(err)
}
