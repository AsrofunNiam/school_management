package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		ID:   "STMN-GURU-01",
		Name: "Sarwidi Utama",
		Role: "guru",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Role, user.Role)
	require.NotZero(t, user.CreatedAt)
}
