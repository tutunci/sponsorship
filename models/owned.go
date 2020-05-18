package models

import (
	"github.com/Kamva/mgm/v2"
)

type Owned struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:",name"`
	Description      string `json:"description" bson:",description"`
	Title            string `json:"title" bson:",title"`
	Status           string `json:"status" bson:",status"`
}

func CreateOwned(name, description, title, status string) *Owned {
	return &Owned{
		Name:        name,
		Description: description,
		Title:       title,
		Status:      status,
	}

}
