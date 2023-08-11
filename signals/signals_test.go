package signals

import (
	"os"
	"syscall"
	"testing"

	"github.com/comtom/portsService/logger"
)

func Test_signalHandler(t *testing.T) {
	type args struct {
		signal os.Signal
		logger *logger.Logger
	}

	logger := logger.NewTestLogger()

	tests := []struct {
		name string
		args args
	}{
		{name: "int", args: args{syscall.SIGINT, logger}},
		{name: "term", args: args{syscall.SIGTERM, logger}},
		{name: "quit", args: args{syscall.SIGQUIT, logger}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signalHandler(tt.args.signal, tt.args.logger)

			// TODO: check resolve action is executed before exit() is called
		})
	}
}
