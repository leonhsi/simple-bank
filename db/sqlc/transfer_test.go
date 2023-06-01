package db

import (
	"context"
	"testing"
	"time"

	"github.com/leonhsi/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	var account1 Account = createRandomAccout(t)
	var account2 Account = createRandomAccout(t)
	CreateRandomTransfer(t, account1, account2)
}

func TestGetTranfer(t *testing.T) {
	var account1 Account = createRandomAccout(t)
	var account2 Account = createRandomAccout(t)
	var transfer Transfer = CreateRandomTransfer(t, account1, account2)

	var getTransfer, err = testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)

	require.Equal(t, transfer.ID, getTransfer.ID)
	require.Equal(t, getTransfer.FromAccountID, account1.ID)
	require.Equal(t, getTransfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, getTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, getTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccout(t)
	account2 := createRandomAccout(t)

	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
	}
}

func TestListTransferErr(t *testing.T) {
	account1 := createRandomAccout(t)
	account2 := createRandomAccout(t)

	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         -1,
		Offset:        -1,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, transfers)
}
