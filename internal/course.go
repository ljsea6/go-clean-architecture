package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid course ID")

// CourseID represents the course unique identifier.
type CourseID struct {
	value string
}

// NewCourseID instantiate the VO for CourseID.
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{
		value: v.String(),
	}, nil
}

// String type converts the CourseID into string.
func (id CourseID) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")

// CourseName represents the course name.
type CourseName struct {
	value string
}

// NewCourseName instantiate VO for CourseName
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	return CourseName{
		value: value,
	}, nil
}

// String type converts the CourseName into string.
func (name CourseName) String() string {
	return name.value
}

var ErrEmptyCourseDuration = errors.New("the field Course Duration can not be empty")

// CourseDuration represents the course duration.
type CourseDuration struct {
	value string
}

// NewCourseDuration instantiate VO for CourseDuration
func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyCourseDuration
	}

	return CourseDuration{
		value: value,
	}, nil
}

// String type converts the CourseDuration into string.
func (duration CourseDuration) String() string {
	return duration.value
}

// CourseRepository defines the expected behaviour from a course storage.
//
//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
	All(ctx context.Context) ([]*Course, error)
}

// Course is the data structure that represents a course.
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course.
func NewCourse(id, name, duration string) (Course, error) {
	idV0, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	nameV0, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationV0, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idV0,
		name:     nameV0,
		duration: durationV0,
	}, nil
}

// ID returns the course unique identifier.
func (c Course) ID() CourseID {
	return c.id
}

// Name returns the course name.
func (c Course) Name() CourseName {
	return c.name
}

// Duration returns the course duration.
func (c Course) Duration() CourseDuration {
	return c.duration
}
