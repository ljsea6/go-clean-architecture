package bootstrap

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ljsea6/go-clean-architecture/internal/platform/server"
	"github.com/ljsea6/go-clean-architecture/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_defaultHost string = "localhost"
	_defaultPort string = "8080"
	dbUser       string = "codely"
	dbPass       string = "codely"
	dbHost       string = "localhost"
	dbPort       string = "3306"
	dbName       string = "codely"
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

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
