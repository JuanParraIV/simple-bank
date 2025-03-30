package db

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

	arg := CreateUserParams{
		Username:       "jotamario",
		HashedPassword: "123456",
		FullName:       "Juan Mario",
		Email:          "jmparra.dev@gmail.com",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.Role)
	require.NotZero(t, user.CreatedAt)
}
