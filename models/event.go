package models

import (
	"time"
)

type Event struct {
	ID      int64     `json:"Id" db:"ID"`
	PetID   int64     `json:"petId" db:"PetID"`
	Date    time.Time `json:"date" validate:"required" `
	Type    string    `json:"type" validate:"required" db:"Type"`
	Remark  string    `json:"remark" db:"Remark"`
	Name    string    `json:"-" db:"Name"`
	DateStr *string   `json:"-" db:"Date"`
}

type PetEvents struct {
	ID    int64   `json:"Id"`
	Name  string  `json:"name"`
	Event []Event `json:"events"`
}
