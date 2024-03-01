package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"testing"

// 	"github.com/aadgraha/school_management/util"
// 	"github.com/stretchr/testify/require"
// )

// func createRandomUser(t *testing.T) User {
// 	arg := CreateUserParams{
// 		ID:   util.RandomUserId(),
// 		Name: util.RandomName(),
// 		Role: util.RandomRole(),
// 	}
// 	user, err := testQueries.CreateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user)
// 	require.Equal(t, arg.ID, user.ID)
// 	require.Equal(t, arg.Name, user.Name)
// 	require.Equal(t, arg.Role, user.Role)
// 	return user
// }
// func TestCreateUser(t *testing.T) {
// 	createRandomUser(t)
// }

// func TestGetUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	user2, err := testQueries.GetUser(context.Background(), user1.ID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user1.Name, user2.Name)
// 	require.Equal(t, user1.Role, user2.Role)
// 	// require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }

// func TestUpdateUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	arg := UpdateUserParams{
// 		ID:   user1.ID,
// 		Name: util.RandomName(),
// 		Role: util.RandomRole(),
// 		ID_2: util.RandomUserId(),
// 	}
// 	user2, err := testQueries.UpdateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, arg.ID_2, user2.ID)
// 	require.Equal(t, arg.Name, user2.Name)
// 	require.Equal(t, arg.Role, user2.Role)
// }

// func TestDeleteUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	err := testQueries.DeleteUser(context.Background(), user1.ID)
// 	require.NoError(t, err)

// 	user2, err := testQueries.GetUser(context.Background(), user1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, user2)
// }

// func TestListUser(t *testing.T) {
// 	// for i := 0; i < 15; i++ {
// 	// 	createRandomUser(t)
// 	// }
// 	arg := ListUsersParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}
// 	users, err := testQueries.ListUsers(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, users, 5)

// 	fmt.Println(users)

// 	for _, user := range users {
// 		require.NotEmpty(t, user)
// 	}

// }
