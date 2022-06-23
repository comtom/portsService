package main

import (
	"os"
	"testing"
)

func Test_initSignals(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases. Couldn't implement because time constraints
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initSignals()
		})
	}
}

func Test_signalHandler(t *testing.T) {
	type args struct {
		signal os.Signal
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases. Couldn't implement because time constraints
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signalHandler(tt.args.signal)
		})
	}
}
