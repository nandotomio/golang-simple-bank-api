package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nandotomio/golang-simple-bank-api/util"
	"github.com/stretchr/testify/require"
)

func mockCreateAccountParams() CreateAccountParams {
	return CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func makeCreateAccount(params CreateAccountParams) (Account, error) {
	return testQueries.CreateAccount(context.Background(), params)
}

func TestCreateAccount(t *testing.T) {
	arg := mockCreateAccountParams()
	account, err := makeCreateAccount(arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	arg := mockCreateAccountParams()
	account, err1 := makeCreateAccount(arg)
	require.NoError(t, err1)
	retrievedAccount, err2 := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, account)

	require.Equal(t, account.ID, retrievedAccount.ID)
	require.Equal(t, account.Owner, retrievedAccount.Owner)
	require.Equal(t, account.Balance, retrievedAccount.Balance)
	require.Equal(t, account.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, retrievedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err2 := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err2)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account, err := makeCreateAccount(mockCreateAccountParams())
	require.NoError(t, err)

	err2 := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err2)

	deletedAccount, err3 := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err3)
	require.EqualError(t, err3, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccount(t *testing.T) {
	owner := util.RandomOwner()
	for i := 0; i < 10; i++ {
		makeCreateAccount(CreateAccountParams{
			Owner:    owner,
			Balance:  util.RandomMoney(),
			Currency: util.RandomCurrency(),
		})
	}

	arg := ListAccountsParams{
		Owner:  owner,
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, owner, account.Owner)
	}
}
