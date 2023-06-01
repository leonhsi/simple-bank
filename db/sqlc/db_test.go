package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithTx(t *testing.T) {
	tx := new(sql.Tx)
	q := New(testDB)

	q = q.WithTx(tx)
	require.Equal(t, q.db, tx)
}
