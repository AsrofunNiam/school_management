package db

import (
	"context"
	"testing"

	"github.com/aadgraha/school_management/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTeacherTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan CreateTeacherResult)
	createUserParams := make(chan CreateUserParams)

	for i := 0; i < n; i++ {

		go func() {
			createUserParam := CreateUserParams{
				ID: util.RandomUserId(), Name: util.RandomName(), Role: "teacher",
			}
			result, err := store.CreateTeacherTx(context.Background(), createUserParam)
			errs <- err
			results <- result
			createUserParams <- createUserParam
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		createUserParam := <-createUserParams
		require.NotEmpty(t, createUserParam)

		user := result.User
		require.NotEmpty(t, user)
		require.Equal(t, createUserParam.ID, user.ID)
		require.Equal(t, createUserParam.Name, user.Name)
		require.Equal(t, createUserParam.Role, user.Role)
		require.NotZero(t, user.CreatedAt)

		teacher := result.Teacher
		require.NotEmpty(t, teacher)
		require.Equal(t, user.ID, teacher.Nip)
		require.NotZero(t, teacher.ID)

		_, err = store.GetUser(context.Background(), user.ID)
		require.NoError(t, err)
		_, err = store.GetTeacher(context.Background(), teacher.ID)
		require.NoError(t, err)
	}

}
