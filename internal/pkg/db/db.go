// Package db provide *gorm.DB for interaction with the database through GORM methods.
package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Internal interface compatible with a logger.Writer.
// Used to configure a custom DB logger.
type Logger interface {
	Printf(format string, args ...any)
}

// Provides *gorm.DB with custom options when creating an object.
type dbSettings struct {
	customLogger    Logger
	logLevel        logger.LogLevel
	translateError  bool
	ignoreNotFound  bool
	disableColorful bool
}

// Type for options for DB struct initializing.
type Option func(*dbSettings)

// Returns new DB instance with connection to given DSN.
// Options can be set with "WithSmth" funcs.
func New(dsn string, options ...Option) (*gorm.DB, error) {
	dbStorage := &dbSettings{
		customLogger:    log.Default(),
		logLevel:        logger.Info,
		translateError:  false,
		ignoreNotFound:  false,
		disableColorful: false,
	}

	// apply all options to customize DB struct
	for _, opt := range options {
		opt(dbStorage)
	}

	gormDB, err := gorm.Open(
		withConn(dsn),
		&gorm.Config{
			// set UTC time zone
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			Logger: logger.New(
				dbStorage.customLogger,
				logger.Config{
					LogLevel:                  dbStorage.logLevel,
					IgnoreRecordNotFoundError: dbStorage.ignoreNotFound,
					Colorful:                  !dbStorage.disableColorful,
				},
			),
			TranslateError: dbStorage.translateError,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("open db connection: %w", err)
	}

	dbStorage.customLogger.Printf("App successfully connected to DB!")
	return gormDB, nil
}

// Set custom logger for DB. Optional.
func WithLogger(customLogger Logger) Option {
	return func(d *dbSettings) {
		d.customLogger = customLogger
	}
}

// Set error log level for DB. Optional.
func WithErrorLogLevel() Option {
	return func(d *dbSettings) {
		d.logLevel = logger.Error
	}
}

// Set warn log level for DB. Optional.
func WithWarnLogLevel() Option {
	return func(d *dbSettings) {
		d.logLevel = logger.Warn
	}
}

// Set translate error parameter true. Optional.
func WithTranslateError() Option {
	return func(d *dbSettings) {
		d.translateError = true
	}
}

// Set ignore record not found error parameter true. Optional.
func WithIgnoreNotFound() Option {
	return func(d *dbSettings) {
		d.ignoreNotFound = true
	}
}

// Set colorful log output false. Optional.
func WithDisableColorful() Option {
	return func(d *dbSettings) {
		d.disableColorful = true
	}
}

// Set connection for DB. Required.
// In this case used SQLite as DB.
func withConn(dsn string) gorm.Dialector {
	return sqlite.Open(dsn)
}
