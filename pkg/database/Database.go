package database

import (
	"errors"
	"fmt"
	. "github.com/GoosvandenBekerom/skateboarding-api/pkg/domain"
)

type Database struct {
	tricks []Trick
}

func GetDatabase() Database {
	return Database{
		tricks: []Trick{
			{
				Name:   "Kickflip",
				Stance: Default,
			},
			{
				Name:   "Kickflip",
				Stance: Fakie,
			},
			{
				Name:   "Kickflip",
				Stance: Nollie,
			},
		},
	}
}

func (db Database) GetAllTricks() []Trick {
	return db.tricks
}

func (db Database) GetTrickByName(name string) (Trick, error) {
	for _, trick := range db.tricks {
		if trick.Name == name {
			return trick, nil
		}
	}
	return Trick{}, errors.New(fmt.Sprintf("No trick found with name %s", name))
}
