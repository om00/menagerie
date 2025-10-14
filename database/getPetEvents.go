package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/om00/menagerie/models"
)

func (db Db) GetPetEvents(id int64) ([]models.Event, error) {

	query := squirrel.Select(
		"p.ID as PetID",
		"p.Name as Name",
		"e.ID",
		"e.Date",
		"e.Type",
		"e.Remark",
	).
		From("pet p").
		LeftJoin("event e ON p.ID = e.PetID").
		Where(squirrel.Eq{"p.ID": id}).
		OrderBy("e.Date DESC")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var events []models.Event
	err = sqlx.Select(db.MainDB, &events, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get pet events: %w", err)
	}

	return events, nil
}
