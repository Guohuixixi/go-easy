// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/Guohuixixi/go-easy/bootstrap/internal"
	"github.com/Guohuixixi/go-easy/core"
)

// Injectors from wire.go:

func InitApp() (core.Application, error) {
	viper := internal.NewViper()
	config := internal.NewConfig(viper)
	db := internal.NewMysql(config)
	cmdable := internal.NewRedis(config)
	logger := internal.NewZap(config)
	engine := internal.NewServer(config)
	application := core.NewApplication(config, viper, db, cmdable, logger, engine)
	return application, nil
}