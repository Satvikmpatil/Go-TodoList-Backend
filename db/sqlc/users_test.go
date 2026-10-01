package db

import (
	"context"
	"testing"

	"github.com/Satvikmpatil/tolist/db/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	username := util.RandomUser()
	user, err := testQueries.CreateUser(context.Background(), username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, username, user.Username)
	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.Username, user2.Username)
}

func TestGetUserByUsername(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.Username, user2.Username)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateRandomUser(t)
	}
	users, err := testQueries.ListUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)
	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.UserID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.UserID)
	require.Error(t, err)
	require.Empty(t, user2)
}
