//go:generate wire
//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func initApp(ConnectionInfo) (*NatsApp, error) {
	wire.Build(
		NewNatsApp,
		NewUserHandler,
		NewDashboardHandler,
		NewUserService,
		NewDashboardService,
		NewUserRepo,
		NewDashboardRepo,
		NewDefaultConfig,
		NewDB,
	)
	return nil, nil // These return values are ignored.
}
