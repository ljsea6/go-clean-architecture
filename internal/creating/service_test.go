package creating

import (
	"context"
	"errors"
	"testing"

	mooc "github.com/ljsea6/go-clean-architecture/internal"
	"github.com/ljsea6/go-clean-architecture/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_CourseService_CreateCourse_RepositoryError(t *testing.T) {
	courseID, courseName, courseDuration := "4f3bd17a-4ae2-406e-9a75-66b831a1aac3", "example", "example duration"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(errors.New("error"))

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Success(t *testing.T) {
	courseID, courseName, courseDuration := "4f3bd17a-4ae2-406e-9a75-66b831a1aac3", "example", "example duration"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
