package migratedb

import (
	"database/sql"
	"path/filepath"

	"github.com/ForAeons/ternary"
	migrate "github.com/rubenv/sql-migrate"
)

func Migrate(db *sql.DB, dir string, step int) (int, error) {
	absPath, err := filepath.Abs("./migrations/")
	if err != nil {
		return 0, err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: absPath,
	}

	n, err := migrate.ExecMax(
		db,
		"postgres",
		migrations,
		ternary.If[migrate.MigrationDirection](dir == "up").
			Then(migrate.Up).
			Else(migrate.Down),
		step,
	)
	return n, err
}
