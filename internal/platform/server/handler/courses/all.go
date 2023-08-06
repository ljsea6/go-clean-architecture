package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/ljsea6/go-clean-architecture/internal"
)

type allResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

func AllHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		courses, err := courseRepository.All(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if len(courses) == 0 {
			ctx.Status(http.StatusNoContent)
			return
		}

		var response []allResponse
		for _, course := range courses {
			response = append(response, allResponse{
				ID:       course.ID().String(),
				Name:     course.Name().String(),
				Duration: course.Duration().String(),
			})
		}

		ctx.JSON(http.StatusOK, response)
	}
}
