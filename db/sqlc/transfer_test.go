package db

import (
	"context"
	"testing"

	"github.com/nandotomio/golang-simple-bank-api/util"
	"github.com/stretchr/testify/require"
)

func mockCreateTransferParams(account1 Account, account2 Account) CreateTransferParams {
	return CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
}

func TestCreateTransfer(t *testing.T) {
	account1, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	arg := mockCreateTransferParams(account1, account2)
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}
