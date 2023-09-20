package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{Owner: util.RandOwner(), Balance: util.RandomMoney(), Currency: util.RandomCurrency()}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err) // checks that the erro is nil and automatically fails if it is not

	require.NotEmpty(t, account) //checks that the account object is not an empty object

	// Check if the created fields match the Account params passed
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)

	// Check that account.ID/account.CreatedAt is available

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateAccout(t *testing.T) {
	// empty context that doesn't have any
	//  associated cancellation signals, deadlines, or request-scoped values
	account1 := createRandomAccount(t)

	updatedAccount := UpdateAccountParams{ID: account1.ID, Balance: util.RandomMoney()}

	account2, err := testQueries.UpdateAccount(context.Background(), updatedAccount)

	require.NoError(t, err)

	require.NotEmpty(t, account2)

	require.Equal(t, updatedAccount.Balance, account2.Balance)

	require.Equal(t, updatedAccount.ID, account2.ID)

	require.Equal(t, account2.Currency, account1.Currency)

	require.Equal(t, account2.Owner, account1.Owner)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, account2)

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)

	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
