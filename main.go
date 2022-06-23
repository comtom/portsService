package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/comtom/portsService/ports"
	"go.uber.org/zap"
)

// this could be extracted as a flag to support more use-cases
const filename = "ports.json"

func signalHandler(signal os.Signal) {
	fmt.Printf("\nCaught signal: %+v", signal)
	fmt.Println("\nWait for 1 second to finish processing")
	time.Sleep(1 * time.Second)

	switch signal {
	// ctrl+c
	case syscall.SIGINT:
		fmt.Println("Signal interrupt triggered.")

	// kill -SIGTERM XXXX [XXXX - PID for your program]
	case syscall.SIGTERM:
		fmt.Println("Signal terminte triggered.")

	// kill -SIGQUIT XXXX [XXXX - PID for your program]
	case syscall.SIGQUIT:
		fmt.Println("Signal quit triggered.")
	}

	// TODO: do something else here (close connections, finish parsing files, finish serving on-flight requests, etc)
	fmt.Println("\nFinished cleaning up. Shutting down")
	os.Exit(0)
}

func initSignals() {
	// test specs asked to handle KILL signal, and it cannot be handled.
	// When a SIGKILL is sent, linux kernel will terminate the process without giving the option to handle it
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	signalHandler(<-signalChanel)
}

func main() {
	// setup logging
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// signal handling
	go initSignals()

	// json is a map, might be good to load only keys first and then objects

	// load file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	portsMap := ports.LoadPorts(f)

	// present the data
	// TODO: would add a REST endpoint to be able to query ports, as having the data in memory and only print isn't that usefull
	// for now we just print basic port data
	if len(portsMap) == 0 {
		panic("no ports loaded! shutting down service")
	}

	logger.Info("loaded ports")
	for k, port := range portsMap {
		logger.Info("port ", zap.String("unilocs", k), zap.String("name", port.Name), zap.String("country", port.Country))
	}
}
