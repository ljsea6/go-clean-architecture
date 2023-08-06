package server

import (
	"fmt"
	"log"

	mooc "github.com/ljsea6/go-clean-architecture/internal"
	"github.com/ljsea6/go-clean-architecture/internal/platform/server/handler/courses"
	"github.com/ljsea6/go-clean-architecture/internal/platform/server/handler/health"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine           *gin.Engine
	httpAddr         string
	courseRepository mooc.CourseRepository
}

func New(host string, port string, courseRepository mooc.CourseRepository) *Server {
	srv := &Server{
		engine:           gin.New(),
		httpAddr:         fmt.Sprintf("%s:%s", host, port),
		courseRepository: courseRepository,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
	s.engine.GET("/courses", courses.AllHandler(s.courseRepository))
}
