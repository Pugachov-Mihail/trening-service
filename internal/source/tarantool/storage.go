package tarantool

import (
	"context"
	"fmt"
	"github.com/tarantool/go-tarantool/v2"
	_ "github.com/tarantool/go-tarantool/v2/datetime"
	_ "github.com/tarantool/go-tarantool/v2/decimal"
	_ "github.com/tarantool/go-tarantool/v2/uuid"
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

	opt := tarantool.Opts{Timeout: time.Second}

	conn, err := tarantool.Connect(ctx, dialer, opt)

	if err != nil {
		log.Error("Unable to connect to database", "error", err)
		return nil, err
	}

	return &TarantoolStorage{db: conn}, nil
}

func (t TarantoolStorage) TreningListSourse(ctx context.Context, page, offset int32, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	defer func() {
		if err := t.db.Close(); err != nil {
			log.Error("Error closing connection", "error", err)
		}
	}()

	res := tarantool.NewSelectRequest("trenning").
		Limit(uint32(page)).
		Offset(uint32(offset)).
		Iterator(tarantool.IterAll)

	resp, err := t.db.Do(res).Get()
	if err != nil {
		log.Error("Error getting data", "error", err)
	}

	var result []trening_v1.GetTreningList
	var name []interface{}

	for _, v := range resp {
		name = append(name, v)
	}
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
