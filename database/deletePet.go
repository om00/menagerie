package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (db Db) DeletePet(id int64) error {
	query := squirrel.Delete("pet").Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	result, err := db.MainDB.Exec(sql, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("pet with ID %d not found", id)
	}

	return nil
}
