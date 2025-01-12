package app

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	trening "trening/internal/api/grpc"
)

type App struct {
	log        *slog.Logger
	grpcServer *grpc.Server
	port       int
	appName    string
}

func New(log *slog.Logger, port int, appName string, treningService trening.Trening) *App {
	grpcServer := grpc.NewServer()

	trening.RegisterTreningServiceApi(grpcServer, treningService)

	reflection.Register(grpcServer)
	return &App{log, grpcServer, port, appName}
}

func (a *App) Run() error {
	log := a.log.With(fmt.Sprintf("%s service Running", a.appName), a.port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		log.Error(fmt.Sprintf("Error Run %s service"), a.appName)
		return fmt.Errorf("%s: %w", fmt.Sprintf("%s Run", a.appName), err)
	}

	log.Info(fmt.Sprintf("%s Run", a.appName), slog.String("addr", l.Addr().String()))

	if err := a.grpcServer.Serve(l); err != nil {
		log.Error(fmt.Sprintf("Error Run GRPC %s service", a.appName))
		return fmt.Errorf("%s: %w", fmt.Sprintf("%s Run", a.appName), err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.With(fmt.Sprintf("%s service Spoping", a.appName), a.port).
		Info(fmt.Sprintf("%s GRPC stop", a.appName))

	a.grpcServer.GracefulStop()
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
