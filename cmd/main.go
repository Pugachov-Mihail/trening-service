package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	apigrpc "trening/internal/app/grpc"
	"trening/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log := config.SetupLoger(cfg.Env)

	log.Debug("Init app")

	application := apigrpc.New(log, cfg.GRPC.Port, *cfg)

	go application.GRPC.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stop app", slog.String("signal", sign.String()))

	application.GRPC.Stop()
	log.Info("app stop")
}
