package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/om00/menagerie/models"
)

func (db Db) CreatePet(pet models.Pet) (int64, error) {
	query, args, err := squirrel.
		Insert("pet").
		Columns("Name", "Owner", "Species", "Birth", "Death").
		Values(pet.Name, pet.Owner, pet.Species, pet.Birth, pet.Death).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to build insert query: %v", err)
	}

	result, err := db.MainDB.Exec(query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to insert pet: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	return id, nil
}
