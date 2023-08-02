package mysql

const (
	sqlCourseTable string = "courses"
)

type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}
