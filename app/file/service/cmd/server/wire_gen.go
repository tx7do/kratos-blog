// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-blog/app/file/service/internal/biz"
	"kratos-blog/app/file/service/internal/conf"
	"kratos-blog/app/file/service/internal/data"
	"kratos-blog/app/file/service/internal/server"
	"kratos-blog/app/file/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewGormClient(confData, logger)
	client := data.NewRedisClient(confData, logger)
	dataData, cleanup, err := data.NewData(db, client, logger)
	if err != nil {
		return nil, nil, err
	}
	attachmentRepo := data.NewAttachmentRepo(dataData, logger)
	attachmentUseCase := biz.NewAttachmentUseCase(attachmentRepo, logger)
	attachmentService := service.NewAttachmentService(logger, attachmentUseCase)
	grpcServer := server.NewGRPCServer(confServer, logger, attachmentService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
