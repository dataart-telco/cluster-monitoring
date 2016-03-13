package thread

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WaitCtrlC() {
	var signal_channel chan os.Signal
	signal_channel = make(chan os.Signal, 2)
	signal.Notify(signal_channel, os.Interrupt, syscall.SIGTERM)
	<-signal_channel
}

func Schedule(step int, what func()) chan int {
	ticker := time.NewTicker(time.Duration(step) * time.Second)
	quit := make(chan int)
	go func() {
		for {
			select {
				case <- quit:
					return
				case <-ticker.C:
					what()
			}
		}
	}()
	return quit
}
