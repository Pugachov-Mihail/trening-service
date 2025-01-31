package apigrpc

import (
	"fmt"
	"log/slog"
	"trening/internal/app"
	"trening/internal/config"
	"trening/internal/service"
	"trening/internal/source/posgres"
	"trening/internal/source/tarantool"
)

type App struct {
	GRPC *app.App
}

func New(log *slog.Logger, port int, cfg config.Config) *App {
	store, err := posgres.NewStorage(cfg.StorageDb, log)
	if err != nil {
		panic("Failed to create store: " + fmt.Sprint(err))
	}

	tstore, err := tarantool.NewTarantoolStorage(cfg.TarantoolDB, log)
	if err != nil {
		panic("Failed to create store: " + fmt.Sprint(err))
	}

	treningService := service.New(log, cfg.GRPC.TokenTTL, cfg, store, tstore)
	trening := app.New(log, port, cfg.App, treningService)

	return &App{GRPC: trening}
}
