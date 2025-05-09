// Package migrate contains migrate interface to control migrations
// and its implementations for different databases.
package migrate

type Migrate interface {
	// Returns version and isDirty value.
	Status() (version uint, isDirty bool, err error)
	// Apply all migrations.
	Up() error
	// Rollback all migrations.
	Down() error
	// Migrate up if n > 0, and down if n < 0.
	Step(n int) error
	// Set a specific migration version.
	Force(n int) error
	// Close connection with database and migrations source.
	Close() error
}
