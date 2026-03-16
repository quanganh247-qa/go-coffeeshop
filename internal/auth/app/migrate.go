//go:build migrate

package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts   = 5
	_defaultTimeout    = time.Second
	_migrationFilePath = "db/migrations"
)

func init() {
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		fmt.Fprintf(os.Stderr, "migrate: environment variable not declared: PG_URL\n")
		return // Don't exit, maybe it's intentional? Or use panic?
	}

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		cur, _ := os.Getwd()
		dir := filepath.Dir(cur + "/../..") // Adjusted from /../../.. based on test

		fmt.Printf("Migrate: looking for migrations at file://%s/%s\n", dir, _migrationFilePath)

		m, err = migrate.New(fmt.Sprintf("file://%s/%s", dir, _migrationFilePath), databaseURL)
		if err == nil {
			break
		}

		fmt.Printf("Migrate: postgres is trying to connect, attempts left: %d\n", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Migrate: postgres connect error: %s\n", err)
		return
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Fprintf(os.Stderr, "Migrate: up error: %s\n", err)
		return
	}

	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("Migrate: no change")
		return
	}

	fmt.Println("Migrate: up success")
}
