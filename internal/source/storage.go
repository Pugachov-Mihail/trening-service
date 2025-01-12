package source

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"
	conf "trening/internal/config"
	trening_v1 "trening/protos/gen/trening.v1"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath conf.Storage, log *slog.Logger) (*Storage, error) {
	//TODO при использовании поправить лог
	log = log.With("connect db", "trening-service")
	connDb := fmt.Sprintf("postgres://" + storagePath.UserDb + ":" + storagePath.PassDb + "@" + storagePath.Host +
		":" + storagePath.PortDb + "/" + storagePath.DbName + "?sslmode=disable")

	db, err := sql.Open("postgres", connDb)

	if err = db.Ping(); err != nil {
		log.Error("Ошибка базы ", err)
		panic("Ошибка подключения к бд")
	}

	return &Storage{db: db}, nil
}

func (s *Storage) TreningListSourse(ctx context.Context, page, offset int32, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	//TODO
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx, "")
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []trening_v1.GetTreningList

	if err = row.Scan(result); err != nil {
		log.Error("Error scanning row", "error", err)
		return nil, err
	}

	return result, nil
}

func (s *Storage) AddTreningSourse(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx, "")
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []trening_v1.GetTreningList

	if err = row.Scan(result); err != nil {
		log.Error("Error scanning row", "error", err)
		return nil, err
	}

	return result, nil
}

func (s *Storage) GetCurrentTreningService(ctx context.Context, idTrening, userId int64, log *slog.Logger) (trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx, "")
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return trening_v1.GetTreningList{}, err
	}

	var result trening_v1.GetTreningList

	if err = row.Scan(result); err != nil {
		log.Error("Error scanning row", "error", err)
		return trening_v1.GetTreningList{}, err
	}

	return result, nil
}
func (s *Storage) DeletedTreningUserService(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx, "")
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []trening_v1.GetTreningList

	if err = row.Scan(result); err != nil {
		log.Error("Error scanning row", "error", err)
		return nil, err
	}

	return result, nil
}

func (s *Storage) GetUserTreningService(ctx context.Context, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx, "")
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []trening_v1.GetTreningList

	if err = row.Scan(result); err != nil {
		log.Error("Error scanning row", "error", err)
		return nil, err
	}

	return result, nil
}
