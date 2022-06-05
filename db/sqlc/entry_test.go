package db

import (
	"context"
	"testing"
	"time"

	"github.com/Aaleena/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	arg := AddEntryParams{
		AccountID: util.RandomAccountID(),
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.AddEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestAddEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntries(t *testing.T) {
	arg := AddEntryParams{
		AccountID: 11,
		Amount:    util.RandomMoney(),
	}
	_, err := testQueries.AddEntry(context.Background(), arg)
	_, err1 := testQueries.AddEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NoError(t, err1)

	entries, err := testQueries.GetEntry(context.Background(), arg.AccountID)

	for _, entry := range entries {
		require.NoError(t, err)
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
		require.NotEmpty(t, entry.Amount)
		require.NotEmpty(t, entry.CreatedAt)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}
	entry1, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)
	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.Equal(t, arg.Amount, entry1.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)
	//require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry1)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 0,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.AccountID)
		require.NotZero(t, entry.CreatedAt)
	}
}
