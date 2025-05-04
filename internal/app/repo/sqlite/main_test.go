package sqlite

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"gorm.io/gorm"

	"BookShelf/internal/pkg/db"
)

const (
	_dbDSN = "file:../../../../db/sqlite3.db?_foreign_keys=on"
)

var (
	_dbStorage *gorm.DB
	_uid       string
)

func TestMain(m *testing.M) {
	var err error

	// open DB connection
	_dbStorage, err = db.New(_dbDSN,
		db.WithTranslateError(),
		db.WithIgnoreNotFound(),
		db.WithDisableColorful(),
	)
	if err != nil {
		log.Fatalf("get db connection: %v", err)
	}

	// create uid based on UNIX-time
	_uid = strconv.FormatInt(time.Now().Unix(), 10)

	log.Printf("UID: %s", _uid)

	// run tests
	os.Exit(m.Run())
}
