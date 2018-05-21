package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// defer function only executed if the program ends
	defer func() {
		log.Println("done")
	}()

	timer := time.NewTimer(time.Second * 5)

	select {
	case <-timer.C: // timer finish
	case <-sigCh: // program interruption
	}
}
