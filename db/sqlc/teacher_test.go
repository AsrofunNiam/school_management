package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/aadgraha/school_management/util"
	"github.com/stretchr/testify/require"
)

func createRandomTeacher(t *testing.T) Teacher {
	user := createRandomUser(t)
	arg := CreateTeacherParams{Nip: util.RandomUserId(), UserID: user.ID}
	teacher, err := testQueries.CreateTeacher(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, teacher)
	require.Equal(t, arg.Nip, teacher.Nip)
	require.WithinDuration(t, user.CreatedAt, teacher.CreatedAt, time.Second)
	return teacher

}
func TestCreateTeacher(t *testing.T) {
	createRandomTeacher(t)
}

func TestListTeacher(t *testing.T) {
	for i := 0; i < 15; i++ {
		createRandomTeacher(t)
	}
	arg := ListTeachersParams{
		Limit:  5,
		Offset: 5,
	}
	teachers, err := testQueries.ListTeachers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, teachers, 5)

	for _, teacher := range teachers {
		require.NotEmpty(t, teacher)
	}
}

func TestGetTeacher(t *testing.T) {
	teacher1 := createRandomTeacher(t)
	teacher2, err := testQueries.GetTeacher(context.Background(), teacher1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, teacher2)

	require.Equal(t, teacher1.ID, teacher2.ID)
	require.Equal(t, teacher1.Nip, teacher2.Nip)
	require.NotEmpty(t, teacher1.CreatedAt)
}

func TestUpdateTeacher(t *testing.T) {
	teacher1 := createRandomTeacher(t)
	require.NotEmpty(t, teacher1)
	argUpdateTeacher := UpdateTeacherParams{Nip: util.RandomUserId(), ID: teacher1.ID}
	teacher2, err := testQueries.UpdateTeacher(context.Background(), argUpdateTeacher)
	require.NoError(t, err)
	require.NotEmpty(t, teacher2)
	require.Equal(t, argUpdateTeacher.ID, teacher2.ID)
	require.Equal(t, argUpdateTeacher.Nip, teacher2.Nip)
}

func TestDeleteTeacher(t *testing.T) {
	teacher1 := createRandomTeacher(t)
	require.NotEmpty(t, teacher1)
	err := testQueries.DeleteTeacher(context.Background(), teacher1.ID)
	require.NoError(t, err)

	teacher2, err := testQueries.GetTeacher(context.Background(), teacher1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, teacher2)
}
