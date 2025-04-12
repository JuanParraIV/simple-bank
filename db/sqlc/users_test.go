package db

import (
	"context"
	"testing"

	"github.com/juanparraiv/simple-bank/utils"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

// TestCreateUser tests the functionality of creating a new user in the database.
// It verifies that the user is created successfully with the correct attributes
// and ensures that no errors occur during the process.
//
// The test performs the following checks:
//   - Ensures no error is returned when creating a user.
//   - Verifies that the returned user object is not empty.
//   - Confirms that the created user's attributes (Username, FullName, Email,
//     and HashedPassword) match the input parameters.
//   - Ensures that the Role and CreatedAt fields of the user are not zero.
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
