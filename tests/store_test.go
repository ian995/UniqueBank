package tests

import (
	"context"
	"fmt"
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

	existed := make(map[int]bool)

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
		_, err = store.GetTransfer(context.Background(), transfer.IDTransfer)
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


		//check account balance
		fromAccount := result.FromAccount 
		require.NoError(t, err)
		require.NotEmpty(t, fromAccount)

		toAccount := result.ToAccount
		require.NoError(t, err)
		require.NotEmpty(t, toAccount)

		diff1 :=  account1.Balance - fromAccount.Balance
		fmt.Println("diff1 >>", diff1,  account1.Balance, fromAccount.Balance)
		diff2 :=   toAccount.Balance - account2.Balance
		fmt.Println("diff2 >>", diff2,  toAccount.Balance, account2.Balance)
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		fmt.Println("k >>", k)
		require.True(t, k>=1 && k<=n)
		require.NotContains(t, existed, k)
		existed[k] = true

	}

	updatedAccount1, err := store.GetAccount(context.Background(), account1.IDAccount)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.IDAccount)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
	
}