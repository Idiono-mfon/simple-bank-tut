package db

import (
	"context"
	"testing"
	"time"

	"github.com/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, AccountID int64) Entry {
	arg := CreateEntryParams{AccountID: AccountID, Amount: util.RandomMoney()}

	// Owner: util.RandOwner(), Balance: util.RandomMoney(), Currency: util.RandomCurrency()

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err) // checks that the erro is nil and automatically fails if it is not

	require.NotEmpty(t, entry) //checks that the account object is not an empty object

	// Check if the created fields match the Account params passed
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	// Check that account.ID/account.CreatedAt is available

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account.ID)

}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)

	entry1 := createRandomEntry(t, account.ID)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err) // checks that the erro is nil and automatically fails if it is not

	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {

	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account.ID)
	}
	arg := ListEntriesParams{
		Limit:     5,
		Offset:    5,
		AccountID: account.ID,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
