package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTeacher(t *testing.T) Teacher {
	user := createRandomUser(t)
	nip := user.ID
	teacher, err := testQueries.CreateTeacher(context.Background(), nip)
	require.NoError(t, err)
	require.NotEmpty(t, teacher)
	require.Equal(t, nip, teacher.Nip)
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
