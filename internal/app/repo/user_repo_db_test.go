package repo

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/pkg/db"
)

const (
	_dbPath = "../../../db/sqlite3.db"
)

var (
	_dbRepo UserRepoDB
	_uid    string
)

func TestMain(m *testing.M) {
	// open DB connection
	dbStorage, err := db.New(
		_dbPath,
		db.WithTranslateError(),
		db.WithIgnoreNotFound(),
		db.WithDisableColorful(),
	)
	if err != nil {
		log.Fatalf("get db connection: %v", err)
	}

	// create user DB repo
	_dbRepo = NewUserRepoDB(dbStorage)
	// create uid based on UNIX-time
	_uid = strconv.FormatInt(time.Now().Unix(), 10)

	log.Printf("UID: %s", _uid)

	// run tests
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Log("Create new user")

	newUser := entity.User{
		Login:    "new_user_" + _uid,
		Password: []byte("123"),
	}

	err := _dbRepo.Create(&newUser)
	assert.NoError(t, err)

	t.Logf("New user: %+v", newUser)
}

func TestCreateDuplicate(t *testing.T) {
	t.Log("Try to create user duplicate")

	newUser := entity.User{
		Login:    "new_user_" + _uid,
		Password: []byte("123"),
	}

	err := _dbRepo.Create(&newUser)
	assert.ErrorIs(t, err, errors.ErrAlreadyExists)
}

func TestGetByLogin(t *testing.T) {
	t.Log("Get user by login")

	existingUser := entity.User{Login: "new_user_" + _uid}

	err := _dbRepo.GetByLogin(&existingUser)
	assert.NoError(t, err)

	t.Logf("Existing user: %+v", existingUser)
}

func TestGetByLoginUnexisting(t *testing.T) {
	t.Log("Try to get user by unexisting login")

	unexistingUser := entity.User{Login: "new_user"}

	err := _dbRepo.GetByLogin(&unexistingUser)
	assert.ErrorIs(t, err, errors.ErrNotFound)
}
