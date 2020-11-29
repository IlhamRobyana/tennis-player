package pg_storage

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var player = &entity.Player{
	ID:       1,
	Username: "Ilham",
	Password: "12345678",
}

func NewMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New() // mock sql.DB
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gdb, err := gorm.Open("postgres", db) // open gorm db
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}
	return gdb, mock
}

func TestLogin(t *testing.T) {
	client, mock := NewMock()
	playerStorage := &Player{client}
	defer client.Close()

	query := `SELECT \* FROM "players" WHERE \(username \= \$1\) ORDER BY "players"\."id" ASC LIMIT 1`

	rows := sqlmock.NewRows([]string{"id", "username", "password"}).
		AddRow(player.ID, player.Username, player.Password)
	mock.ExpectQuery(query).WithArgs(player.Username).WillReturnRows(rows)

	correctPlayer, err := playerStorage.GetByUsername(player.Username)
	assert.NotNil(t, correctPlayer)
	assert.NoError(t, err)
}

func TestLoginError(t *testing.T) {
	client, mock := NewMock()
	playerStorage := &Player{client}
	defer client.Close()

	query := `SELECT \* FROM "players" WHERE \(username \= \$1\) ORDER BY "players"\."id" ASC LIMIT 1`
	mock.ExpectQuery(query).WithArgs(player.Username)

	_, err := playerStorage.GetByUsername(player.Username)
	assert.Error(t, err)
}
