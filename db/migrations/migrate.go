package migrations

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate() error {
	// Initialize the migration
	m, err := migrate.New("file://db/migrations/", "sqlite3://employees.db")
	if err != nil {
		return fmt.Errorf("could not initialize migration: %w", err)
	}

	defer func() {
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			fmt.Printf("error closing migration: srcErr=%v, dbErr=%v\n", srcErr, dbErr)
		}
	}()

	// Run the migration
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	} else if err == migrate.ErrNoChange {
		fmt.Println("No new migrations to apply.")
	}

	return nil
}
