package database

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/om00/menagerie/models"
)

func (db Db) UpdatePet(id int64, updateReq models.UpdatePetReq) (int64, error) {

	updateBuilder := squirrel.Update("pet")

	if updateReq.Name != "" {
		updateBuilder = updateBuilder.Set("Name", updateReq.Name)
	}
	if updateReq.Owner != "" {
		updateBuilder = updateBuilder.Set("Owner", updateReq.Owner)
	}
	if updateReq.Species != "" {
		updateBuilder = updateBuilder.Set("Species", updateReq.Species)
	}
	if updateReq.Birth != nil {
		updateBuilder = updateBuilder.Set("Birth", *updateReq.Birth)
	}
	if updateReq.Death != nil {
		updateBuilder = updateBuilder.Set("Death", *updateReq.Death)
	}

	updateBuilder = updateBuilder.Where(squirrel.Eq{"ID": id})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build update query: %v", err)
	}

	result, err := db.MainDB.Exec(query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to update pet: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("no pet found with ID %d", id)
	}

	return rowsAffected, nil
}
