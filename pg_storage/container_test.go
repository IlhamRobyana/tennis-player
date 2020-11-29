package pg_storage

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/stretchr/testify/assert"
)

var container = &entity.Container{
	ID:       1,
	PlayerID: 1,
	Capacity: 3,
	Balls:    0,
}

func TestPutBall(t *testing.T) {
	client, mock := NewMock()
	containerStorage := &Container{client}
	defer client.Close()

	query := `SELECT \* FROM "containers" WHERE \(player_id\=\$1 AND balls \< capacity\)`

	rows := sqlmock.NewRows([]string{"id", "player_id", "capacity", "balls"}).
		AddRow(container.ID, container.PlayerID, container.Capacity, container.Balls)
	mock.ExpectQuery(query).WithArgs(container.PlayerID).WillReturnRows(rows)

	mock.ExpectBegin()

	query = `UPDATE "containers" SET "balls" \= \$1, "capacity" \= \$2, "id" \= \$3, "player_id" \= \$4  WHERE "containers"\."id" \= \$5 AND \(\(id\=\$6\)\)`
	mock.ExpectExec(query).WithArgs(container.Balls+1, container.Capacity, container.ID, container.PlayerID, container.ID, container.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	correctContainer, err := containerStorage.PutBall(container.PlayerID)
	assert.NotNil(t, correctContainer)
	assert.NoError(t, err)
}
