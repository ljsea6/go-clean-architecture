package mysql

import (
	"context"
	"errors"
	"testing"

	mooc "github.com/ljsea6/go-clean-architecture/internal"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CourseRepository_Save_RepositoryError(t *testing.T) {
	courseID, courseName, courseDuration := "4f3bd17a-4ae2-406e-9a75-66b831a1aac3", "example", "example"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnError(errors.New("error"))

	repo := NewCourseRepository(db)
	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CourseRepository_Save_Success(t *testing.T) {
	courseID, courseName, courseDuration := "4f3bd17a-4ae2-406e-9a75-66b831a1aac3", "example", "example"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db)
	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

func Test_CourseRepository_All_RepositoryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)

	sqlMock.ExpectQuery("SELECT id, name, duration FROM courses").
		WillReturnError(errors.New("error"))

	repo := NewCourseRepository(db)
	_, err = repo.All(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CourseRepository_All_ValidationError(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)

	coursesMockRows := sqlmock.NewRows([]string{"id", "name", "duration"}).
		AddRow("1", "test", "test")

	sqlMock.ExpectQuery("SELECT id, name, duration FROM courses").
		WillReturnRows(coursesMockRows)

	repo := NewCourseRepository(db)
	_, err = repo.All(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CourseRepository_All_Success(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)

	coursesMockRows := sqlmock.NewRows([]string{"id", "name", "duration"}).
		AddRow("bea1b24d-0627-4ea0-aa2b-8af4c6c2a41c", "test", "test")

	sqlMock.ExpectQuery("SELECT id, name, duration FROM courses").
		WillReturnRows(coursesMockRows)

	repo := NewCourseRepository(db)
	_, err = repo.All(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
