package models

import (
	"time"
)

type Pet struct {
	ID      int64      `db:"ID"`
	Name    string     `json:"name"  validate:"required" db:"Name"`
	Owner   string     `json:"owner" validate:"required" db:"Owner"`
	Species string     `json:"species" validate:"required" db:"Species"`
	Birth   *time.Time `json:"birth,omitempty" `
	Death   *time.Time `json:"death,omitempty" `

	BirthStr *string `db:"Birth" json:"-"`
	DeathStr *string `db:"Death" json:"-"`
}

type UpdatePetReq struct {
	Name    string     `json:"name" db:"Name"`
	Owner   string     `json:"owner"  db:"Owner"`
	Species string     `json:"species"  db:"Species"`
	Birth   *time.Time `json:"birth," db:"Birth"`
	Death   *time.Time `json:"death"  db:"Death"`
}

type GetFilter struct {
	ID      int64
	Species string
}
