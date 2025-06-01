package tests

import (
	"context"
	"testing"

	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := repo.NewStore(testDb)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)


	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan repo.TransferTxResult)


	for range n {
		go func ()  {
			result, err := store.TransferTx(context.Background(), repo.TransferTxParam{
				FromIDAccount: account1.IDAccount,
				ToIDAccount:   account2.IDAccount,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	for range n {
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.Equal(t, account1.IDAccount, transfer.FromIDAccount)
		require.Equal(t, account2.IDAccount, transfer.ToIDAccount)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.IDTransfer)
		require.NotZero(t, transfer.CreateAt)		
		_, err = store.GetTransfer(context.Background(), account1.IDAccount)
		require.NoError(t, err)


		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.IDAccount, fromEntry.IDAccount)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.IDEntries)
		require.NotZero(t, fromEntry.CreateAt)
		_, err = store.GetEntry(context.Background(), fromEntry.IDEntries)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.IDAccount, toEntry.IDAccount)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.IDEntries)
		require.NotZero(t, toEntry.CreateAt)
		_, err = store.GetEntry(context.Background(), toEntry.IDEntries)
		require.NoError(t, err)


	}
	
}