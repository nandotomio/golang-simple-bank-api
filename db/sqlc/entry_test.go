package db

import (
	"context"
	"testing"

	"github.com/nandotomio/golang-simple-bank-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
}
