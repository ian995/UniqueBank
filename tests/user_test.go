package tests

import (
	"context"
	"testing"
	"time"

	"github.com/ian995/UniqueBank/internal/repo"
	"github.com/ian995/UniqueBank/pkg/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) repo.User {
	arg := repo.CreateUserParams{
		Username: 	 utils.RandomOwner(),
		HashedPassword: utils.RandomString(10),
		FullName:     utils.RandomOwner(),
		Email:        utils.RandomEmail(),

	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.HashedPassword)
	require.True(t, user.PasswordChangeAt.IsZero())
	require.NotZero(t, user.CreateAt)
	return user
}
func TestUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangeAt, user2.PasswordChangeAt, time.Second)
	require.WithinDuration(t, user1.CreateAt, user2.CreateAt, time.Second)

}
