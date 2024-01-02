package di

import (
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/app"
	"github.com/eganow/partners/sampler/api/v1/features/auth/business_logic/app/repository"
	"github.com/eganow/partners/sampler/api/v1/features/auth/pkg"
)

type AuthInjector struct {
	UseCase *app.AuthUseCase // the use case
	Repo    pkg.Repository   // the repository
}

// NewAuthInjector creates a new AuthInjector instance.
func NewAuthInjector() *AuthInjector {
	injector := &AuthInjector{}

	// create the repository
	injector.Repo = repository.NewNoopAuthRepository()

	// create the use case
	injector.UseCase = app.NewAuthUseCase(injector.Repo)

	return injector
}
