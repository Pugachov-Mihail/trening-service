package main

import (
	"context"
	"github.com/robfig/cron/v3"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	apigrpc "trening/internal/app/grpc"
	"trening/internal/config"
	"trening/internal/source/posgres"
	"trening/internal/source/tarantool"
	"trening/internal/sourse_sync"
)

func main() {
	cfg := config.MustLoad()
	log := config.SetupLoger(cfg.Env)
	syncCron := cron.New()
	var wg sync.WaitGroup

	log.Debug("Init app")

	application := apigrpc.New(log, cfg.GRPC.Port, *cfg)

	go application.GRPC.MustRun()

	ctx := context.Background()

	taskSync := sourse_sync.NewSync(syncCron, log)
	storageTSync, err := tarantool.NewTarantoolStorage(cfg.TarantoolDB, log)
	storageSSync, err := posgres.NewStorage(cfg.StorageDb, log)
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	taskSync.Start(storageTSync, storageSSync, ctx)

	syncCron.Start()
	wg.Wait()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stop app", slog.String("signal", sign.String()))

	application.GRPC.Stop()
	log.Info("app stop")
}
