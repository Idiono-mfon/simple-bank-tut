package db

import (
	"context"
	"testing"
	"time"

	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, FromAccountID int64, ToAccountID int64) Transfer {
	arg := CreateTransferParams{FromAccountID: FromAccountID, ToAccountID: ToAccountID, Amount: util.RandomMoney()}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err) // checks that the erro is nil and automatically fails if it is not

	require.NotEmpty(t, transfer) //checks that the account object is not an empty object

	// Check if the created fields match the Account params passed
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	// Check that account.ID/account.CreatedAt is available

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)

	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1.ID, account2.ID)

}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)

	account2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, account1.ID, account2.ID)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err) // checks that the erro is nil and automatically fails if it is not

	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transfer.Amount, transfer.Amount)

	require.WithinDuration(t, transfer.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)

}

func TestListTransfers(t *testing.T) {

	account1 := createRandomAccount(t)

	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1.ID, account2.ID)
	}
	arg := ListTransfersParams{
		Limit:         5,
		Offset:        5,
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
