//go:generate wire
//go:build wireinject
// +build wireinject

package internal

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
		NewHttpCleint,
		wire.Bind(new(HTTP), new(*HttpCleint)),
	)
	return nil, nil // These return values are ignored.
}
