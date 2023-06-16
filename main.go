package main

import (
	"log"
	"sync"
)

func main() {
	app := InitializeApp()
	log.Println("App is initialized")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(a App) {
		defer wg.Done()
		a.httpServer.StartHttpServer()
	}(*app)
	wg.Wait()
}
