package source

import (
	"encoding/json"
	"fmt"
	"log/slog"
	trening_v1 "trening/protos/gen/trening.v1"
)

type Trenings struct {
	Id           int
	Titile       string
	Descriptions string
	Image        string
	Price        float64
	mdate        string
	FirstName    string
	LastName     string
	Exercise     []Exercise
	Data         []uint8 `json:"exercise"`
}

type Exercise struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Avatar      string      `json:"avatar"`
	Video       string      `json:"video"`
	Description string      `json:"description"`
	Approach    []*Approach `json:"approaches"`
}

type Approach struct {
	Weigth float32 `json:"weigth"`
	Repeat int32   `json:"repeat"`
	Rest   int64   `json:"rest"`
}

func (t *Trenings) UnmarshalExercise() error {
	var e []Exercise

	if err := json.Unmarshal(t.Data, &e); err != nil {
		return err
	}

	t.Exercise = append(t.Exercise, e...)

	return nil
}

func (t *Trenings) ReturnsResult(logger *slog.Logger) trening_v1.GetTreningList {
	var exercise []*trening_v1.Exercise
	var approach []*trening_v1.Approach

	for _, ex := range t.Exercise {
		for _, a := range ex.Approach {
			approach = append(approach, &trening_v1.Approach{
				Weigth: a.Weigth,
				Repeat: a.Repeat,
				Rest:   a.Rest,
			})
		}
		logger.Info("Approach data:", "approach", fmt.Sprintf("%+v", approach))

		exercise = append(exercise,
			&trening_v1.Exercise{
				Title:       ex.Title,
				Avatar:      ex.Avatar,
				Video:       ex.Video,
				Description: ex.Description,
				Approach:    approach,
			},
		)
	}

	return trening_v1.GetTreningList{
		Id:          int64(t.Id),
		Title:       t.Titile,
		Description: t.Descriptions,
		Image:       t.Image,
		TrenerInfo: &trening_v1.Trener{
			LastName: t.LastName,
			Name:     t.FirstName,
		},
		TrenningInfo: exercise,
	}
}
