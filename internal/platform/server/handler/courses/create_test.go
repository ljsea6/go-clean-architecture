package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ljsea6/go-clean-architecture/internal/platform/storage/storagemocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(courseRepository))

	t.Run("given an invalid request return 400", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:   "77a8edbc-9c0c-4388-b227-471f506a9b3d",
			Name: "Demo Course",
		}

		payload, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(payload))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		courseRepository.AssertNumberOfCalls(t, "Save", 0)
	})

	t.Run("given an invalid id return 400", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:   "7",
			Name: "Demo Course",
		}

		payload, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(payload))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		courseRepository.AssertNumberOfCalls(t, "Save", 0)
	})

	t.Run("given a valid request return 201", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "77a8edbc-9c0c-4388-b227-471f506a9b3d",
			Name:     "Demo Course",
			Duration: "10 months",
		}

		payload, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(payload))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		courseRepository.AssertNumberOfCalls(t, "Save", 1)
	})
}
