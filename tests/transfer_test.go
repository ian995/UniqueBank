package tests

import (
	"context"
	"testing"
	"time"

	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/ian995/UniqueBank/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateRandomTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := repo.CreateTransferParams{
		FromIDAccount: account1.IDAccount,
		ToIDAccount:   account2.IDAccount,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromIDAccount, transfer.FromIDAccount)
	require.Equal(t, arg.ToIDAccount, transfer.ToIDAccount)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.IDTransfer)
	require.NotZero(t, transfer.CreateAt)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := repo.CreateTransferParams{
		FromIDAccount: account1.IDAccount,
		ToIDAccount:   account2.IDAccount,
		Amount:        utils.RandomMoney(),
	}

	transfer1, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer1)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.IDTransfer)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.IDTransfer, transfer2.IDTransfer)
	require.Equal(t, transfer1.FromIDAccount, transfer2.FromIDAccount)
	require.Equal(t, transfer1.ToIDAccount, transfer2.ToIDAccount)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreateAt, transfer2.CreateAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := repo.CreateTransferParams{
			FromIDAccount: account1.IDAccount,
			ToIDAccount:   account2.IDAccount,
			Amount:        utils.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)
	}

	listArg := repo.ListTransfersParams{
		Limit:  5,
		Offset: 0,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), listArg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.NotZero(t, transfer.IDTransfer)
		require.NotZero(t, transfer.FromIDAccount)
		require.NotZero(t, transfer.ToIDAccount)
		require.NotZero(t, transfer.Amount)
		require.NotZero(t, transfer.CreateAt)
	}
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.NotZero(t, transfer.IDTransfer)
		require.NotZero(t, transfer.FromIDAccount)
		require.NotZero(t, transfer.ToIDAccount)
		require.NotZero(t, transfer.Amount)
		require.NotZero(t, transfer.CreateAt)
	}
}
func TestListTransferByToAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := repo.CreateTransferParams{
			FromIDAccount: account1.IDAccount,
			ToIDAccount:   account2.IDAccount,
			Amount:        utils.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)
	}

	listArg := repo.ListTransfersByToAccountParams{
		ToIDAccount: account2.IDAccount,
		Limit:     5,
		Offset:    0,
	}
	transfers, err := testQueries.ListTransfersByToAccount(context.Background(), listArg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, account2.IDAccount, transfer.ToIDAccount)
	}
}

func TestListTransferByFromAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := repo.CreateTransferParams{
			FromIDAccount: account1.IDAccount,
			ToIDAccount:   account2.IDAccount,
			Amount:        utils.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)
	}

	listArg := repo.ListTransfersByFromAccountParams{
		FromIDAccount: account1.IDAccount,
		Limit:     5,
		Offset:    0,
	}
	transfers, err := testQueries.ListTransfersByFromAccount(context.Background(), listArg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.IDAccount, transfer.FromIDAccount)
	}
}