package utils

import (
	"os"
	"os/signal"
	"syscall"
)

var globalWaitChannel chan os.Signal

// SignalQuit used for testing
func SignalQuit() {
	if globalWaitChannel != nil {
		globalWaitChannel <- os.Interrupt
	}
}

// WaitSignal waits for a syscall.X
func WaitSignal() os.Signal {
	globalWaitChannel = make(chan os.Signal)
	signal.Notify(
		globalWaitChannel,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-globalWaitChannel
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}
