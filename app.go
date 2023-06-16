package main

import "github.com/andybeak/hexagonal-demo/internal/handlers"

func ProvideApp(
	httpServer *handlers.HttpService,
) *App {
	return &App{
		httpServer: *httpServer,
	}
}

type App struct {
	httpServer handlers.HttpService
}
