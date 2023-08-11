package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/comtom/portsService/logger"
	"go.uber.org/zap"
)

const exitTimeout = 1

func signalHandler(signal os.Signal, logger *logger.Logger) {
	logger.Info("received a system signal", zap.String("signal", signal.String()))
	logger.Info(fmt.Sprintf("waiting %v second to finish processing", exitTimeout))
	time.Sleep(exitTimeout * time.Second)

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
	logger.Info("finished cleaning up. shutting down")
	os.Exit(0)
}

func InitSignals(logger *logger.Logger) {
	// test specs asked to handle KILL signal, and it cannot be handled.
	// When a SIGKILL is sent, linux kernel will terminate the process without giving the option to handle it
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	signalHandler(<-signalChanel, logger)
}
