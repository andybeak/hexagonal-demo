//go:build wireinject

package main

import (
	"github.com/andybeak/hexagonal-demo/helpers"
	"github.com/andybeak/hexagonal-demo/internal/core/usecases"
	"github.com/andybeak/hexagonal-demo/internal/handlers"
	"github.com/andybeak/hexagonal-demo/internal/repositories"
	"github.com/google/wire"
)

func InitializeApp() *App {
	wire.Build(
		ProvideApp,
		adapters.ProvideUserRepository,
		usecases.ProvideUserUseCase,
		helpers.ProvideDatabaseParameters,
		helpers.ProvideDatabase,
		handlers.ProvideHttpService,
		handlers.ProvideHttpServiceParams,
		handlers.ProvideRouter,
		handlers.ProvideUserHttpHandler,
	)
	return &App{}
}
