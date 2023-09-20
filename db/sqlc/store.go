package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store instance
func NewStore(db *sql.DB) *Store {

	return &Store{
		db: db,
		// Queries: , // This creates and returns a query object

		Queries: New(db),
	}
}

// Define a method on Store to execute and generate Database transaction

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Create a new queries object for the transaction
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	// Then call the callback function with the created queries object
	// New Query object which also uses the transaction object
	q := New(tx)

	err = fn(q)

	if err != nil {
		// rollback error: rbErr
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}

		return err
	}

	// Everything goes well, commit the transaction
	return tx.Commit()

	// And finally commit or rollback the transaction
	// based on the error returned by that function.

}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"` //when the struct is serialized it json, this field should be called from_account_id
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResults struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

// TransferTx performs a money tansfer from one account to the other
// It creates a transfer record, add account entries and update accounts' vbalance within a single dtbase transaction.

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {

	var result TransferTxResults

	/**
		Closure is often used when we want to get the result from a callback function,
		// Callback function here is a closure
	**/

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		txName := ctx.Value(txKey)

		// Since arg has similar fields with CreateTransferParams struct, we can use this signature to convert to automatically

		// CreateTransferParams{
		// 	FromAccountID: arg.FromAccountID,
		// 	ToAccountID:   arg.ToAccountID,
		// 	Amount:        arg.Amount,
		// }
		// ||
		// ||
		//\||/
		// CreateTransferParams

		fmt.Println(txName, "Create transfer")

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))

		if err != nil {
			return nil
		}

		fmt.Println(txName, "Create entry 1")

		// Money is leaving the account
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.FromAccountID, Amount: -arg.Amount})

		if err != nil {
			return err
		}

		// money is entry the account

		fmt.Println(txName, "Create entry 2")

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{AccountID: arg.ToAccountID, Amount: arg.Amount})

		if err != nil {
			return err
		}

		// fmt.Println(txName, "Get account 1 ")

		// account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)

		// if err != nil {
		// 	return err
		// }

		// fmt.Println(txName, "Update account 1 ")

		// result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{ID: arg.FromAccountID, Balance: account1.Balance - arg.Amount})

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{ID: arg.FromAccountID, Amount: -arg.Amount})

		if err != nil {
			return err
		}

		// fmt.Println(txName, "Get account 2 ")

		// account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)

		// if err != nil {
		// 	return err
		// }

		// fmt.Println(txName, "Update account 2 ")

		// result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{ID: arg.ToAccountID, Balance: account2.Balance + arg.Amount})
		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{ID: arg.ToAccountID, Amount: arg.Amount})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
