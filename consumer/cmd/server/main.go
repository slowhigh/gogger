package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Run() error
	Stop()
}

func main() {
	var server Server

	server, err := InitServer()
	if err != nil {
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	server.Stop()

	slog.Info("Terminated.")
}
