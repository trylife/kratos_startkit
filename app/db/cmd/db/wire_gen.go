// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"startkit/app/db/internal/conf"
	"startkit/app/db/internal/server"
	"startkit/app/db/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	migrationService := service.NewMigrationService()
	grpcServer := server.NewGRPCServer(confServer, migrationService, logger)
	httpServer := server.NewHTTPServer(confServer, migrationService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
