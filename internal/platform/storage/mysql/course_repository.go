package mysql

import (
	"context"
	"database/sql"
	"fmt"

	mooc "github.com/ljsea6/go-clean-architecture/internal"

	"github.com/huandu/go-sqlbuilder"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) mooc.CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

func (r *CourseRepository) All(ctx context.Context) ([]*mooc.Course, error) {
	query, args := sqlbuilder.NewSelectBuilder().
		Select("id", "name", "duration").
		From(sqlCourseTable).
		Build()

	var courses []*mooc.Course

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return courses, fmt.Errorf("error trying to get courses from database: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var id, name, duration string

		if err := rows.Scan(&id, &name, &duration); err != nil {
			return courses, fmt.Errorf("error trying to scan courses from database: %v", err)
		}

		course, err := mooc.NewCourse(id, name, duration)
		if err != nil {
			return courses, err
		}

		courses = append(courses, &course)
	}

	return courses, nil
}
