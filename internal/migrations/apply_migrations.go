package migrations

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

func ApplyMigrations(db *sql.DB) error {
	// Get all migration files
	migrationDir := "internal/migrations"
	files, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %v", err)
	}

	// Sort migrations by name
	var migrationFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrationFiles = append(migrationFiles, filepath.Join(migrationDir, file.Name()))
		}
	}
	sort.Strings(migrationFiles)

	// Execute each migration
	for _, migrationFile := range migrationFiles {
		content, err := ioutil.ReadFile(migrationFile)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", migrationFile, err)
		}

		// Execute SQL statements
		statements := strings.Split(string(content), "--")
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt != "" {
				// Eliminar comentarios de SQL
				stmt = strings.ReplaceAll(stmt, "--", "-- ")
				stmt = strings.TrimSpace(stmt)
				if stmt != "" {
					_, err := db.Exec(stmt)
					if err != nil {
						return fmt.Errorf("failed to execute migration %s: %v", migrationFile, err)
					}
					log.Printf("Successfully executed migration %s", migrationFile)
				}
			}
		}
	}

	return nil
}
