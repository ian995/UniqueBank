package tests

import (
	"context"
	"testing"
	"time"

	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/ian995/UniqueBank/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) repo.Account {
	user := createRandomUser(t)

	arg := repo.CreateAccountParams{
		Owner:   user.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.IDAccount)
	require.NotZero(t, account.CreateAt)

	return account
}
func TestAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.IDAccount)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.IDAccount, account2.IDAccount)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreateAt, account2.CreateAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := repo.UpdateAccountParams{
		IDAccount: account1.IDAccount,
		Balance:   utils.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.IDAccount, account2.IDAccount)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreateAt, account2.CreateAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.IDAccount)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.IDAccount)
	require.Error(t, err)
	require.Empty(t, account2)
}


func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := repo.ListAccountsParams{
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