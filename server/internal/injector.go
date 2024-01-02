package internal

import (
	"errors"
	"github.com/eganow/partners/sampler/api/v1/features/auth/di"
)

var (
	// AuthInjector represents a dependency injector for the auth feature.
	AuthInjector *di.AuthInjector
)

// InitializeDependencies initializes the dependencies for all the features.
func InitializeDependencies() error {
	var err error

	// register dependencies
	if AuthInjector = di.NewAuthInjector(); AuthInjector == nil {
		err = errors.New("failed to initialize auth feature dependencies")
	}
	// @todo: register other dependencies here

	return err
}
