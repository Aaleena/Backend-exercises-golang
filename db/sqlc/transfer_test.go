package db

import (
	"context"
	"testing"
	"time"

	"github.com/Aaleena/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	arg := AddTransferEntryParams{
		FromAccountID: util.RandomAccountID(),
		ToAccountID:   util.RandomAccountID(),
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.AddTransferEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestAddTransferEntry(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	transfer1, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer1)

	require.Equal(t, transfer.ID, transfer1.ID)
	require.Equal(t, transfer.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, transfer.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer.Amount, transfer1.Amount)
	require.WithinDuration(t, transfer.CreatedAt, transfer1.CreatedAt, time.Second)
}

func TestGetTransfers(t *testing.T) {
	argAdd := AddTransferEntryParams{
		FromAccountID: 12,
		ToAccountID:   15,
		Amount:        util.RandomMoney(),
	}
	_, err := testQueries.AddTransferEntry(context.Background(), argAdd)
	_, err1 := testQueries.AddTransferEntry(context.Background(), argAdd)

	require.NoError(t, err)
	require.NoError(t, err1)

	argGet := GetTransfersParams{
		FromAccountID: argAdd.FromAccountID,
		ToAccountID:   argAdd.ToAccountID,
	}
	transfers, err := testQueries.GetTransfers(context.Background(), argGet)

	for _, transfer := range transfers {
		require.NoError(t, err)
		require.NotEmpty(t, transfer)
		require.Equal(t, argGet.FromAccountID, transfer.FromAccountID)
		require.Equal(t, argGet.ToAccountID, transfer.ToAccountID)
		require.NotEmpty(t, transfer.CreatedAt)
		require.NotEmpty(t, transfer.ID)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	arg := UpdateTransfersParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}
	transfer1, err := testQueries.UpdateTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer1)
	require.Equal(t, transfer.ID, transfer1.ID)
	require.Equal(t, transfer.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, arg.Amount, transfer1.Amount)
	require.WithinDuration(t, transfer.CreatedAt, transfer1.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	transfer1, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, "sql: no rows in result set")
	require.Empty(t, transfer1)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  3,
		Offset: 1,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 3)

	for _, transfer := range transfers {
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.FromAccountID)
		require.NotZero(t, transfer.ToAccountID)
		require.NotZero(t, transfer.CreatedAt)
	}
}
