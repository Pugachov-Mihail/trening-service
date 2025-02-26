package posgres

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"time"
	conf "trening/internal/config"
	"trening/internal/source"
	trening_v1 "trening/protos/gen/trening.v1"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath conf.Storage, log *slog.Logger) (*Storage, error) {
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

func (s *Storage) TreningListSourse(ctx context.Context, page, offset int32, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := s.db.QueryContext(dbCtx,
		`select * from trening_dm.Get_Programm(_limit := $1, _offset := $2 );`,
		page, offset)

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
		var result source.Trenings

		if err = rows.Scan(&result.Id, &result.Titile,
			&result.Descriptions, &result.Image,
			&result.Price, &result.FirstName,
			&result.LastName, &result.Data); err != nil {
			log.Error("Error scanning row", "error", err)
			return nil, fmt.Errorf("scan error: %w", err)
		}

		if err := result.UnmarshalExercise(); err != nil {
			log.Error("Error unmarshaling exercise", "error", err)
		}

		res = append(res, result.ReturnsResult())
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

	row, err := s.db.QueryContext(dbCtx,
		`SELECT * FROM trening_dm.Add_Current_Trening_For_User(_programm_id := $1, _user_id := $2);`,
		idTrening, userId,
	)

	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []trening_v1.GetTreningList

	for row.Next() {
		var res source.Trenings

		if err = row.Scan(&res.Id, &res.Titile,
			&res.Descriptions, &res.Image, &res.FirstName,
			&res.LastName, &res.Data); err != nil {
			log.Error("Error scanning row", "error", err)
			return nil, err
		}

		if err := res.UnmarshalExercise(); err != nil {
			log.Error("Error unmarshaling exercise", "error", err)
		}

		result = append(result, res.ReturnsResult())
	}

	if err = row.Err(); err != nil {
		log.Error("Error scanning row", "error", err)
	}

	return result, nil
}

func (s *Storage) GetCurrentTreningService(ctx context.Context, idTrening, userId int64, log *slog.Logger) (trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx,
		`SELECT * FROM  trening_dm.Get_Current_Trening(_programm_id := $1)`,
		idTrening)
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return trening_v1.GetTreningList{}, err
	}

	var res source.Trenings

	for row.Next() {
		if err = row.Scan(&res.Id, &res.Titile,
			&res.Descriptions, &res.Image, &res.FirstName,
			&res.LastName, &res.Data); err != nil {
			log.Error("Error scanning row", "error", err)
			return trening_v1.GetTreningList{}, err
		}
		if err := res.UnmarshalExercise(); err != nil {
			log.Error("Error unmarshaling exercise", "error", err)
		}
	}

	return res.ReturnsResult(), nil
}

func (s *Storage) DeletedTreningUserService(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx,
		`SELECT * FROM trening_dm.Delete_Programm_User(_programm_id := $1, _user_id := $2);`,
		idTrening, userId,
	)
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []trening_v1.GetTreningList

	for row.Next() {
		var res source.Trenings

		if err = row.Scan(&res.Id, &res.Titile,
			&res.Descriptions, &res.Image, &res.Data); err != nil {
			log.Error("Error scanning row", "error", err)
			return nil, err
		}

		if err := res.UnmarshalExercise(); err != nil {
			log.Error("Error unmarshaling exercise", "error", err)
		}

		result = append(result, res.ReturnsResult())

	}

	if err = row.Err(); err != nil {
		log.Error("Error scanning row", "error", err)
	}

	return result, nil
}

func (s *Storage) GetUserTreningService(ctx context.Context, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	row, err := s.db.QueryContext(dbCtx,
		`select * from trening_dm.Get_User_Trenings(_user_id := $1);`,
		userId)
	defer func() {
		if err := row.Close(); err != nil {
			log.Error("Error closing connect", "error", err)
		}
	}()

	if err != nil {
		return nil, err
	}

	var result []trening_v1.GetTreningList

	for row.Next() {
		var res source.Trenings

		if err = row.Scan(&res.Id, &res.Titile,
			&res.Descriptions, &res.Image, &res.Data); err != nil {
			log.Error("Error scanning row", "error", err)
			return nil, err
		}

		if err := res.UnmarshalExercise(); err != nil {
			log.Error("Error unmarshaling exercise", "error", err)
		}

		result = append(result, res.ReturnsResult())

	}

	if err = row.Err(); err != nil {
		log.Error("Error scanning row", "error", err)
	}

	return result, nil
}

func (s *Storage) DataForSync(ctx context.Context, log *slog.Logger) ([]source.Trenings, error) {
	var result []source.Trenings

	dbCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	rows, err := s.db.QueryContext(dbCtx, `select * from trening_dm.Get_Programm(_limit := $1, _offset := $2 );`, 1000, 0)

	if err != nil {
		log.Warn("Error during iteration Storage", "error", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Error("Error closing rows", "error", err)
		}
	}()

	for rows.Next() {
		//Дополнить proto trening_v1.GetTreningList
		var res source.Trenings

		if err = rows.Scan(&res.Id, &res.Titile,
			&res.Descriptions, &res.Image,
			&res.Price, &res.FirstName,
			&res.LastName, &res.Data); err != nil {
			log.Error("Error scanning row", "error", err)
		}

		if err = res.UnmarshalExercise(); err != nil {
			log.Error("Error unmarshaling exercise", "error", err)
		}

		result = append(result, res)
	}

	if err := rows.Err(); err != nil {
		log.Error("Error during iteration", "error", err)
	}

	return result, nil
}
