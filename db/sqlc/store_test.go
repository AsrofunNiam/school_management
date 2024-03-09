package db

import (
	"context"
	"testing"
	"time"

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
				Name: util.RandomName(), Role: "teacher",
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
		require.Equal(t, createUserParam.Name, user.Name)
		require.Equal(t, createUserParam.Role, user.Role)
		require.NotZero(t, user.CreatedAt)

		teacher := result.Teacher
		require.NotEmpty(t, teacher)
		require.NotZero(t, teacher.ID)
		require.NotEmpty(t, teacher.CreatedAt)
		require.NotEmpty(t, teacher.UserID)
		require.WithinDuration(t, user.CreatedAt, teacher.CreatedAt, time.Second)

		_, err = store.GetUser(context.Background(), user.ID)
		require.NoError(t, err)
		_, err = store.GetTeacher(context.Background(), teacher.ID)
		require.NoError(t, err)
	}

}

func TestCreateStudentTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan CreateStudentResult)
	createUserParams := make(chan CreateUserParams)

	for i := 0; i < n; i++ {

		go func() {
			createUserParam := CreateUserParams{
				Name: util.RandomName(), Role: "student",
			}
			result, err := store.CreateStudentTx(context.Background(), createUserParam)
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
		require.Equal(t, createUserParam.Name, user.Name)
		require.Equal(t, createUserParam.Role, user.Role)
		require.NotZero(t, user.CreatedAt)

		student := result.Student
		require.NotEmpty(t, student)
		require.NotZero(t, student.ID)
		require.NotEmpty(t, student.CreatedAt)
		require.NotEmpty(t, student.UserID)
		require.WithinDuration(t, user.CreatedAt, student.CreatedAt, time.Second)

		_, err = store.GetUser(context.Background(), user.ID)
		require.NoError(t, err)
		_, err = store.GetStudent(context.Background(), student.ID)
		require.NoError(t, err)
	}

}
