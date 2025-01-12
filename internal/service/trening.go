package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"trening/internal/config"
	trening_v1 "trening/protos/gen/trening.v1"
)

type Trening struct {
	Log      *slog.Logger
	TokenTTl time.Duration
	Cfg      config.Config
	Store    TreningStorage
}

type TreningStorage interface {
	TreningListSourse(ctx context.Context, page, offset int32, log *slog.Logger) ([]trening_v1.GetTreningList, error)
	AddTreningSourse(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error)
	GetCurrentTreningService(ctx context.Context, idTrening, userId int64, log *slog.Logger) (trening_v1.GetTreningList, error)
	DeletedTreningUserService(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error)
	GetUserTreningService(ctx context.Context, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error)
}

// New сервисный слой Trening
func New(log *slog.Logger, ttl time.Duration, cfg config.Config, store TreningStorage) *Trening {
	return &Trening{Log: log, TokenTTl: ttl, Cfg: cfg, Store: store}
}

func (t *Trening) TreningsListService(ctx context.Context, page, offset int32) ([]*trening_v1.GetTreningList, error) {
	log := t.Log.With("trening-list", "trening-list service")

	value, err := t.Store.TreningListSourse(ctx, page, offset, log)
	if err != nil {
		log.Error(fmt.Sprintf("trening list get store: %s", err))
		return nil, fmt.Errorf("treningList: %w", err)
	}

	var treningList []*trening_v1.GetTreningList

	for _, v := range value {
		treningList = append(treningList, &v)
	}

	log.Info("trening list return result")

	return treningList, nil
}

func (t *Trening) AddTreningsUserService(ctx context.Context, idTrening, userId int64) ([]*trening_v1.GetTreningList, error) {
	log := t.Log.With("trening-add", "add-trening-list service")

	value, err := t.Store.AddTreningSourse(ctx, idTrening, userId, log)
	if err != nil {
		log.Error(fmt.Sprintf("trening add get store: %s", err))
		return nil, fmt.Errorf("addTrening: %w", err)
	}

	var treningList []*trening_v1.GetTreningList

	for _, v := range value {
		treningList = append(treningList, &v)
	}

	log.Info("add trening user in list return result")

	return treningList, nil
}

func (t *Trening) GetUserTreningService(ctx context.Context, userId int64) ([]*trening_v1.GetTreningList, error) {
	log := t.Log.With("user-trening", "get user trening service")

	value, err := t.Store.GetUserTreningService(ctx, userId, log)
	if err != nil {
		log.Error(fmt.Sprintf("trening get user trening service: %s", err))
		return nil, fmt.Errorf("getUserTreningService: %w", err)
	}
	var treningList []*trening_v1.GetTreningList

	for _, v := range value {
		treningList = append(treningList, &v)
	}

	log.Info("get-user-trening service return result")

	return treningList, nil
}

func (t *Trening) DeletedTreningUserService(ctx context.Context, idTrening, userId int64) ([]*trening_v1.GetTreningList, error) {
	log := t.Log.With("user-trening", "deleted user trening service")

	value, err := t.Store.DeletedTreningUserService(ctx, idTrening, userId, log)
	if err != nil {
		log.Error(fmt.Sprintf("trening deleted user trening service: %s", err))
	}
	var treningList []*trening_v1.GetTreningList

	for _, v := range value {
		treningList = append(treningList, &v)
	}

	log.Info("delete-user-trening service return result")

	return treningList, nil
}

func (t *Trening) GetCurrentTreningService(ctx context.Context, idTrening, id int64) (trening_v1.GetTreningList, error) {
	log := t.Log.With("user-trening", "get current trening service")

	value, err := t.Store.GetCurrentTreningService(ctx, idTrening, id, log)
	if err != nil {
		log.Error(fmt.Sprintf("trening get current trening service: %s", err))
		return trening_v1.GetTreningList{}, err
	}

	log.Info("delete-user-trening service return result")

	return trening_v1.GetTreningList{
		Id:           value.Id,
		Title:        value.Title,
		Description:  value.Description,
		Image:        value.Image,
		TrenningInfo: value.TrenningInfo,
		TrenerInfo:   value.TrenerInfo,
		Raiting:      value.Raiting,
	}, nil
}
