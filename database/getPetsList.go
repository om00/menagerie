package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/om00/menagerie/models"
)

func (db Db) GetPetsList(filter models.GetFilter) ([]models.Pet, error) {
	query := squirrel.Select(
		"ID",
		"Name",
		"Owner",
		"Species",
		"Birth",
		"Death",
	).From("pet")

	if filter.Species != "" {
		query = query.Where(squirrel.Eq{"Species": filter.Species})
	}

	if filter.ID != 0 {
		query = query.Where(squirrel.Eq{"ID": filter.ID})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %w", err)
	}

	var pets []models.Pet
	err = sqlx.Select(db.MainDB, &pets, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query pets: %w", err)
	}

	return pets, nil
}
