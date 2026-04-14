package db

import (
	"TicketRservation/env"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp() {
	// Build MySQL URL
	mydbs := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		env.DbUser.GetValue(),
		env.DbPassword.GetValue(),
		env.DbHost.GetValue(),
		env.DbPort.GetValue(),
		env.DbName.GetValue(),
	)

	// Build absolute path to migrations
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	sourceurl := "file://" + filepath.Join(cwd, "db", "migrations")

	m, err := migrate.New(sourceurl, mydbs)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	fmt.Println("Migrations applied successfully!")
}
