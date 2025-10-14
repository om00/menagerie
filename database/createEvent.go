package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/om00/menagerie/models"
)

func (db Db) CreateEvent(event models.Event) (int64, error) {
	insertBuilder := squirrel.Insert("event").
		Columns("PetID", "Date", "Type")

	values := []interface{}{event.PetID, event.Date, event.Type}

	if event.Remark != "" {
		insertBuilder = insertBuilder.Columns("Remark")
		values = append(values, event.Remark)
	}

	insertBuilder = insertBuilder.Values(values...)

	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	res, err := db.MainDB.Exec(sql, args...)
	if err != nil {

		return 0, fmt.Errorf("failed to insert event: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	return id, nil
}
