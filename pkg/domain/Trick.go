package domain

import (
	"encoding/json"
)

type Trick struct {
	Name   string
	Stance Stance
}

func (t Trick) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name   string
		Stance string
	}{
		Name:   t.Name,
		Stance: t.Stance.String(),
	})
}
