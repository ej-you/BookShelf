package migrate

import (
	"errors"
	"fmt"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite" // sqlite engine for migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"     // engine for migration files
)

var _ Migrate = (*sqliteMigrate)(nil)

// Migrate implementation.
type sqliteMigrate struct {
	mgrt *gomigrate.Migrate
}

func NewSQLiteMigrate(sourceURL, databaseURL string) (Migrate, error) {
	mgrt, err := gomigrate.New(sourceURL, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create migrate manager: %w", err)
	}
	return &sqliteMigrate{mgrt: mgrt}, nil
}

func (s *sqliteMigrate) Status() (version uint, isDirty bool, err error) {
	v, d, err := s.mgrt.Version()
	if err != nil && !errors.Is(err, gomigrate.ErrNilVersion) {
		return 0, false, fmt.Errorf("migrate status: %w", err)
	}
	return v, d, nil
}

func (s *sqliteMigrate) Up() error {
	err := s.mgrt.Up()
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

func (s *sqliteMigrate) Down() error {
	err := s.mgrt.Down()
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migrate down: %w", err)
	}
	return nil
}

func (s *sqliteMigrate) Step(n int) error {
	err := s.mgrt.Steps(n)
	if err != nil {
		return fmt.Errorf("migrate step: %w", err)
	}
	return nil
}

func (s *sqliteMigrate) Force(n int) error {
	err := s.mgrt.Force(n)
	if err != nil {
		return fmt.Errorf("migrate force: %w", err)
	}
	return nil
}

func (s *sqliteMigrate) Close() error {
	sourceCloseErr, dbCloseErr := s.mgrt.Close()
	if sourceCloseErr != nil && dbCloseErr != nil {
		return fmt.Errorf("close database: %s && close source: %w", dbCloseErr.Error(), sourceCloseErr)
	}
	if sourceCloseErr != nil {
		return fmt.Errorf("close source: %w", sourceCloseErr)
	}
	if dbCloseErr != nil {
		return fmt.Errorf("close database: %w", dbCloseErr)
	}
	return nil
}
