package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	MainDB *sqlx.DB
}

var dbPath string

func CreateConnection() (Db, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)
	dbPath = "mysql://" + dsn
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Print("Error connecting to DB:", err)
	}

	return Db{MainDB: db}, nil
}

func RunMigrations() {
	migrationsDir := "file://migrations"

	m, err := migrate.New(
		migrationsDir,
		dbPath,
	)
	if err != nil {
		log.Fatalf("Could not initialize migration: %v", err)
	}

	defer m.Close()
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	if err == migrate.ErrNoChange {
		fmt.Println("No migrations to run. Database up-to-date.")
	} else {
		fmt.Println("Migrations applied successfully.")
	}
}

func (db Db) IsPetExist(id int64) (bool, error) {
	query := squirrel.Select("1").From("pet").Where(squirrel.Eq{"id": id})

	rawSql, args, err := query.ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to build query: %w", err)
	}

	var exists bool
	err = db.MainDB.QueryRow(rawSql, args...).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, fmt.Errorf("failed to check pet existence: %w", err)
	}

	return exists, nil
}
