package db

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const _dbPath = "../../../db/sqlite3.db"

func TestDBWithoutOptions(t *testing.T) {
	t.Log("Create DB instance with default options")

	dbConn, err := New(_dbPath)
	assert.NoError(t, err, "get db connection")

	t.Logf("DB instance: %+v", dbConn)
}

func TestDBWithCustomLogger(t *testing.T) {
	t.Log("Create DB instance with custom logger")

	customLogger := log.New(os.Stderr, "[LOGGER]\t", log.Ldate|log.Ltime)

	dbConn, err := New(_dbPath, WithLogger(customLogger))
	assert.NoError(t, err, "get db connection")

	t.Logf("DB instance: %+v", dbConn)
}

func TestDBWithAllOptions(t *testing.T) {
	t.Log("Create DB instance with all custom options")

	customLogger := log.New(os.Stdout, "[ALL OPTIONS]\t", log.Ltime)

	dbConn, err := New(
		_dbPath,
		WithLogger(customLogger),
		WithErrorLogLevel(),
		WithTranslateError(),
		WithIgnoreNotFound(),
		WithDisableColorful(),
	)
	assert.NoError(t, err, "get db connection")

	t.Logf("DB instance: %+v", dbConn)
}
