package db

import (
	"context"
	"testing"
	"time"

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

func makeCreateTransfer(params CreateTransferParams) (Transfer, error) {
	return testQueries.CreateTransfer(context.Background(), params)
}

func TestCreateTransfer(t *testing.T) {
	account1, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	arg := mockCreateTransferParams(account1, account2)
	transfer, err := makeCreateTransfer(arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	account1, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	transfer, err := makeCreateTransfer(mockCreateTransferParams(account1, account2))
	require.NoError(t, err)

	retrievedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedTransfer)

	require.Equal(t, transfer.ID, retrievedTransfer.ID)
	require.Equal(t, transfer.FromAccountID, retrievedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, retrievedTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, retrievedTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, retrievedTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	account2, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	for i := 0; i < 5; i++ {
		_, err := makeCreateTransfer(mockCreateTransferParams(account1, account2))
		require.NoError(t, err)
		_, err = makeCreateTransfer(mockCreateTransferParams(account2, account1))
		require.NoError(t, err)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}
}
