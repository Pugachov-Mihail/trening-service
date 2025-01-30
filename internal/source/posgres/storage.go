package posgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
		":" + storagePath.PortDb + "/" + storagePath.DbName + "?sslmode=disable") //&search_path=trening_dm

	log.Info("Connecting to ", connDb)

	db, err := sql.Open("postgres", connDb)

	if err = db.Ping(); err != nil {
		log.Error("Ошибка базы ", err)
		panic(fmt.Sprintf("Ошибка подключения к бд: %v", err))
	}

	return &Storage{db: db}, nil
}

// TODO переписать возвращаемый результат
func (s *Storage) TreningListSourse(ctx context.Context, page, offset int32, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := s.db.QueryContext(dbCtx, `select * from trening_dm.GetProgramm(_limit := $1, _offset := $2 );`, page, offset)

	if err != nil {
		return nil, fmt.Errorf("query context error: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error("Error closing rows", "error", err)
		}
	}()

	var res []trening_v1.GetTreningList

	for rows.Next() {
		//Дополнить proto trening_v1.GetTreningList
		var result struct {
			Id           int
			Titile       string
			Descriptions string
			Image        string
			Price        float64
			FirstName    string
			LastName     string
		}
		if err = rows.Scan(&result.Id, &result.Titile,
			&result.Descriptions, &result.Image,
			&result.Price, &result.FirstName,
			&result.LastName); err != nil {
			log.Error("Error scanning row", "error", err)
			return nil, fmt.Errorf("scan error: %w", err)
		}

		res = append(res, trening_v1.GetTreningList{
			Id:          int64(result.Id),
			Title:       result.Titile,
			Description: result.Descriptions,
			Image:       result.Image,
			TrenerInfo: &trening_v1.Trener{
				LastName: result.LastName,
				Name:     result.FirstName,
			},
		})
	}

	if err := rows.Err(); err != nil {
		log.Error("Error during iteration", "error", err)
		return nil, fmt.Errorf("rows.err: %w", err)
	}

	return res, nil
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
