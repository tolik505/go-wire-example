//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"wire/internal"
	mock "wire/internal/mocks"

	"github.com/google/wire"
)

func initUserHandler(
	internal.ConnectionInfo,
	*mock.MockHTTP,
) (*internal.UserHandler, error) {
	wire.Build(
		internal.NewUserHandler,
		internal.NewUserService,
		internal.NewUserRepo,
		internal.NewDefaultConfig,
		internal.NewDB,
		wire.Bind(new(internal.HTTP), new(*mock.MockHTTP)),
	)
	return nil, nil // These return values are ignored.
}
