// Code generated by mockery v2.32.2. DO NOT EDIT.

package storagemocks

import (
	context "context"

	mooc "github.com/ljsea6/go-clean-architecture/internal"
	mock "github.com/stretchr/testify/mock"
)

// CourseRepository is an autogenerated mock type for the CourseRepository type
type CourseRepository struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *CourseRepository) All(ctx context.Context) ([]*mooc.Course, error) {
	ret := _m.Called(ctx)

	var r0 []*mooc.Course
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*mooc.Course, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*mooc.Course); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*mooc.Course)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, course
func (_m *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mooc.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCourseRepository creates a new instance of CourseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCourseRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CourseRepository {
	mock := &CourseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
