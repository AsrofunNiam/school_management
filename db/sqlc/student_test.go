package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/aadgraha/school_management/util"
	"github.com/stretchr/testify/require"
)

func createRandomStudent(t *testing.T) Student {
	user := createRandomUser(t)
	arg := CreateStudentParams{Nis: util.RandomUserId(), UserID: user.ID}
	student, err := testQueries.CreateStudent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, student)
	require.Equal(t, arg.Nis, student.Nis)
	require.WithinDuration(t, user.CreatedAt, student.CreatedAt, time.Second)
	return student

}
func TestCreateStudent(t *testing.T) {
	createRandomStudent(t)
}

func TestListStudent(t *testing.T) {
	for i := 0; i < 15; i++ {
		createRandomStudent(t)
	}
	arg := ListStudentsParams{
		Limit:  5,
		Offset: 5,
	}
	students, err := testQueries.ListStudents(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, students, 5)

	for _, student := range students {
		require.NotEmpty(t, student)
	}
}

func TestGetStudent(t *testing.T) {
	student1 := createRandomStudent(t)
	student2, err := testQueries.GetStudent(context.Background(), student1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, student2)

	require.Equal(t, student1.ID, student2.ID)
	require.Equal(t, student1.Nis, student2.Nis)
	require.NotEmpty(t, student1.CreatedAt)
}

func TestUpdateStudent(t *testing.T) {
	student1 := createRandomStudent(t)
	require.NotEmpty(t, student1)
	argUpdateStudent := UpdateStudentParams{Nis: util.RandomUserId(), ID: student1.ID}
	student2, err := testQueries.UpdateStudent(context.Background(), argUpdateStudent)
	require.NoError(t, err)
	require.NotEmpty(t, student2)
	require.Equal(t, argUpdateStudent.ID, student2.ID)
	require.Equal(t, argUpdateStudent.Nis, student2.Nis)
}

func TestDeleteStudent(t *testing.T) {
	student1 := createRandomStudent(t)
	require.NotEmpty(t, student1)
	err := testQueries.DeleteStudent(context.Background(), student1.ID)
	require.NoError(t, err)

	student2, err := testQueries.GetStudent(context.Background(), student1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, student2)
}