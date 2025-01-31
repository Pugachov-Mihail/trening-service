package sourse_sync

import (
	"context"
	"github.com/robfig/cron/v3"
	"log/slog"
	"trening/internal/source/posgres"
	"trening/internal/source/tarantool"
)

type Sync struct {
	cron *cron.Cron
	log  *slog.Logger
}

func NewSync(cron *cron.Cron, log *slog.Logger) *Sync {
	return &Sync{cron: cron, log: log}
}

func (s *Sync) Start(tarantool *tarantool.TarantoolStorage, storage *posgres.Storage, ctx context.Context) {
	log := s.log.With("worker", "sync")

	_, err := s.cron.AddFunc("@every 30s", func() {
		go tarantool.SyncData(storage, log, ctx)
	})
	if err != nil {
		log.Error("Error", err)
	}

}

func (s *Sync) Sync() {

}
