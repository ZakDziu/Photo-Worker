package main

import (
	"main/internal/handlers"
	"main/utils"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	bus := utils.StartMQWorker()
	utils.CreateWorkDirectory()
	wg.Add(2)
	go bus.Receiver.Listen(&wg)
	go utils.Exit(bus, &wg)
	http.ListenAndServe(":8080", handlers.StartServer(bus))
	wg.Wait()
}
