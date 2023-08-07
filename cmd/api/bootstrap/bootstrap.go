package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ljsea6/go-clean-architecture/internal/creating"
	"github.com/ljsea6/go-clean-architecture/internal/platform/bus/inmemory"
	"github.com/ljsea6/go-clean-architecture/internal/platform/server"
	"github.com/ljsea6/go-clean-architecture/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_defaultHost    string = "localhost"
	_defaultPort    string = "8080"
	shutdownTimeout        = 10 * time.Second

	dbUser string = "codely"
	dbPass string = "codely"
	dbHost string = "localhost"
	dbPort string = "3306"
	dbName string = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = _defaultHost
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = _defaultPort
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
