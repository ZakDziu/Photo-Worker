package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Exit(bus *Bus, group *sync.WaitGroup) {
	defer group.Done()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		bus.Conn.Close()
		bus.Ch.Close()
		err := os.RemoveAll("./photos/")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("  Close All Connections ")
		os.Exit(1)
	}()
}
