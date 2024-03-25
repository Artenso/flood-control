package service_provider

import (
	"context"

	"github.com/Artenso/FloodControl/internal/service"
)

type App struct {
	serviceProvider *serviceProvider
	service         *service.Service
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initService,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initService(ctx context.Context) error {
	a.service = a.serviceProvider.getService(ctx)
	return nil
}
