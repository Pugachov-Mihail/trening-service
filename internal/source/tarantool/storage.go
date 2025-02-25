package tarantool

import (
	"context"
	"fmt"
	"github.com/tarantool/go-tarantool/v2"
	_ "github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/decimal"
	_ "github.com/tarantool/go-tarantool/v2/uuid"
	"trening/internal/source/posgres"

	"log/slog"
	"time"
	"trening/internal/config"
	trening_v1 "trening/protos/gen/trening.v1"
)

type TarantoolStorage struct {
	db *tarantool.Connection
}

func NewTarantoolStorage(config config.Tarantool, logger *slog.Logger) (*TarantoolStorage, error) {
	log := logger.With("TorantoolStorage", "store")

	ctx, cancelCTX := context.WithTimeout(context.Background(), time.Second)
	defer cancelCTX()

	dialer := tarantool.NetDialer{
		Address:  fmt.Sprintf("%s:%s", config.HostDb, config.PortDb),
		User:     config.UserDb,
		Password: config.PassDb,
	}

	opt := tarantool.Opts{}

	conn, err := tarantool.Connect(ctx, dialer, opt)

	if err != nil {
		log.Error("Unable to connect to database", "error", err)
		return nil, err
	}

	return &TarantoolStorage{db: conn}, nil
}

func (t TarantoolStorage) TreningListSourse(ctx context.Context, page, offset int32, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	type resStruct struct {
		Id           int
		Titile       string
		Descriptions string
		Image        string
		Price        float64
		FirstName    string
		LastName     string
	}
	// Создаем запрос на выборку данных
	res := tarantool.NewSelectRequest("trenning").
		Limit(uint32(page)).
		Offset(uint32(offset))

	data := []resStruct{}

	// Выполняем запрос
	err := t.db.Do(res).GetTyped(&data)
	if err != nil {
		log.Error("Error getting data", "error", err)
		return nil, err
	}

	var result []trening_v1.GetTreningList

	for _, v := range data {
		result = append(result, trening_v1.GetTreningList{
			Id:          int64(v.Id),
			Title:       v.Titile,
			Description: v.Descriptions,
			Image:       v.Image,
			TrenerInfo: &trening_v1.Trener{
				LastName: v.LastName,
				Name:     v.FirstName,
			},
		})
	}

	log.Debug("tarantool get result")

	return result, nil
}

func (t TarantoolStorage) AddTreningSourse(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	//TODO implement me
	panic("implement me")
}

func (t TarantoolStorage) GetCurrentTreningService(ctx context.Context, idTrening, userId int64, log *slog.Logger) (trening_v1.GetTreningList, error) {
	//TODO implement me
	panic("implement me")
}

func (t TarantoolStorage) DeletedTreningUserService(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	//TODO implement me
	panic("implement me")
}

func (t TarantoolStorage) GetUserTreningService(ctx context.Context, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TarantoolStorage) SyncData(storage *posgres.Storage, logger *slog.Logger, ctx context.Context) {
	log := logger.With("TorantoolStorage", "sync")

	res, err := storage.DataForSync(ctx, log)

	req := tarantool.NewCallRequest("add_trennings_if_not_exist").
		Args([]interface{}{res})

	_, err = t.db.Do(req).Get()
	if err != nil {
		log.Warn("Error getting data", "error", err)
		return
	}

	log.Info("TorantoolStorage sync complete")
}
