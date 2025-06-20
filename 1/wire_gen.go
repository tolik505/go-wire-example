// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func initApp(connectionInfo ConnectionInfo) (*NatsApp, error) {
	config := NewDefaultConfig()
	db, err := NewDB(connectionInfo)
	if err != nil {
		return nil, err
	}
	userRepo, err := NewUserRepo(config, db)
	if err != nil {
		return nil, err
	}
	userService := NewUserService(userRepo)
	userHandler := NewUserHandler(userService)
	natsApp, err := NewNatsApp(userHandler)
	if err != nil {
		return nil, err
	}
	return natsApp, nil
}
