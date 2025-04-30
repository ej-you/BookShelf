package sqlite

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"BookShelf/internal/app/entity"
	"BookShelf/internal/app/errors"
	"BookShelf/internal/app/repo"
	"BookShelf/internal/pkg/db"
)

const (
	_dbPath = "../../../../db/sqlite3.db"
)

var (
	_dbRepo repo.UserRepoDB
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
	require.NoError(t, err)

	t.Logf("New user: %+v", newUser)
}

func TestCreateDuplicate(t *testing.T) {
	t.Log("Try to create user duplicate")

	newUser := entity.User{
		Login:    "new_user_" + _uid,
		Password: []byte("123"),
	}

	err := _dbRepo.Create(&newUser)
	require.ErrorIs(t, err, errors.ErrAlreadyExists)
}

func TestGetByLogin(t *testing.T) {
	t.Log("Get user by login")

	existingUser := entity.User{Login: "new_user_" + _uid}

	err := _dbRepo.GetByLogin(&existingUser)
	require.NoError(t, err)

	t.Logf("Existing user: %+v", existingUser)
}

func TestGetByLoginUnexisting(t *testing.T) {
	t.Log("Try to get user by unexisting login")

	unexistingUser := entity.User{Login: "new_user"}

	err := _dbRepo.GetByLogin(&unexistingUser)
	require.ErrorIs(t, err, errors.ErrNotFound)
}
