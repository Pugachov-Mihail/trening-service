package source

import (
	"encoding/json"
	trening_v1 "trening/protos/gen/trening.v1"
)

type Trenings struct {
	Id           int     `json:"id"`
	Titile       string  `json:"titile"`
	Descriptions string  `json:"description"`
	Image        string  `json:"image"`
	Price        float64 `json:"price,omitempty"`
	Mdate        string  `json:"mdate"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
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

func (t *Trenings) ReturnsResult() trening_v1.GetTreningList {
	var exercise []*trening_v1.Exercise
	var approach []*trening_v1.Approach

	for i, ex := range t.Exercise {
		for _, a := range t.Exercise[i].Approach {
			approach = append(approach, &trening_v1.Approach{
				Weigth: a.Weigth,
				Repeat: a.Repeat,
				Rest:   a.Rest,
			})
		}

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
