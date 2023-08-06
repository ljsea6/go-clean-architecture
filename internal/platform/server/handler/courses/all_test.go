package courses

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mooc "github.com/ljsea6/go-clean-architecture/internal"
	"github.com/ljsea6/go-clean-architecture/internal/platform/storage/storagemocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_ALl_Error(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("All", mock.Anything).Return(nil, errors.New("error"))

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/courses", AllHandler(courseRepository))

	t.Run("return a internal server error 500", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		courseRepository.AssertNumberOfCalls(t, "All", 1)
	})
}

func TestHandler_ALl_No_Content(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("All", mock.Anything).Return(nil, nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/courses", AllHandler(courseRepository))

	t.Run("given a valid request return 204", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusNoContent, res.StatusCode)
		courseRepository.AssertNumberOfCalls(t, "All", 1)
	})
}

func TestHandler_ALl_OK(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/courses", AllHandler(courseRepository))

	t.Run("given a valid request return 200", func(t *testing.T) {
		course, err := mooc.NewCourse("4f3bd17a-4ae2-406e-9a75-66b831a1aac3", "example", "example duration")
		require.NoError(t, err)

		courses := []*mooc.Course{
			&course,
		}
		courseRepository.On("All", mock.Anything).Return(courses, nil)

		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
		courseRepository.AssertNumberOfCalls(t, "All", 1)
	})
}
