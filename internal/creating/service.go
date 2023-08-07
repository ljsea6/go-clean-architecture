package creating

import (
	"context"

	mooc "github.com/ljsea6/go-clean-architecture/internal"
)

type CourseService struct {
	courseRepository mooc.CourseRepository
}

func (s *CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	return s.courseRepository.Save(ctx, course)
}

func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}
