package tests

import (
	"context"
	"testing"
	"time"

	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/ian995/UniqueBank/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateRandomEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := repo.CreateEntryParams{
		IDAccount: account.IDAccount,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.IDAccount, entry.IDAccount)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.IDEntries)
	require.NotZero(t, entry.CreateAt)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := repo.CreateEntryParams{
		IDAccount: account.IDAccount,
		Amount:    utils.RandomMoney(),
	}

	entry1, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.IDEntries)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.IDEntries, entry2.IDEntries)
	require.Equal(t, entry1.IDAccount, entry2.IDAccount)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreateAt, entry2.CreateAt, time.Second)
}
func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := repo.CreateEntryParams{
			IDAccount: account.IDAccount,
			Amount:    utils.RandomMoney(),
		}

		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg := repo.ListEntriesParams{
		Limit:     5,
		Offset:    0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestListEntriesByAccount(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := repo.CreateEntryParams{
			IDAccount: account.IDAccount,
			Amount:    utils.RandomMoney(),
		}

		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
	}

	arg := repo.ListEntriesByAccountParams{
		IDAccount: account.IDAccount,
		Limit:     5,
		Offset:    0,
	}

	entries, err := testQueries.ListEntriesByAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}