package db

import (
	"context"
	"testing"
	"time"

	"github.com/nandotomio/golang-simple-bank-api/util"
	"github.com/stretchr/testify/require"
)

func mockCreateEntryParams(account Account) CreateEntryParams {
	return CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
}

func makeCreateEntry(params CreateEntryParams) (Entry, error) {
	return testQueries.CreateEntry(context.Background(), params)
}

func TestCreateEntry(t *testing.T) {
	account, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)

	arg := mockCreateEntryParams(account)
	entry, err := makeCreateEntry(arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	account, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)

	arg := mockCreateEntryParams(account)
	entry, err := makeCreateEntry(arg)
	require.NoError(t, err)

	retrievedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedEntry)

	require.Equal(t, entry.ID, retrievedEntry.ID)
	require.Equal(t, entry.AccountID, retrievedEntry.AccountID)
	require.Equal(t, entry.Amount, retrievedEntry.Amount)
	require.WithinDuration(t, entry.CreatedAt, retrievedEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		_, err := makeCreateEntry(mockCreateEntryParams(account))
		require.NoError(t, err)
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
