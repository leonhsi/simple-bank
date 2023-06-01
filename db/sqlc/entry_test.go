package db

import (
	"context"
	"testing"
	"time"

	"github.com/leonhsi/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccout(t)
	createRandomEntries(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccout(t)
	entry1 := createRandomEntries(t, account)

	entry2, err := testQueries.GetEntries(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccout(t)
	for i := 0; i < 10; i++ {
		createRandomEntries(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}

func TestListEntriesErr(t *testing.T) {
	arg := ListEntriesParams{
		AccountID: 0,
		Limit:     -1,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, entries)
}
