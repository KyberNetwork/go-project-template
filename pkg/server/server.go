package server

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server to serve the service
type Server struct {
	s        *gin.Engine
	bindAddr string
	l        *zap.SugaredLogger
}

// New server
func New(bindAddr string) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	s := &Server{
		s:        engine,
		bindAddr: bindAddr,
		l:        zap.S(),
	}
	gin.SetMode(gin.ReleaseMode)
	s.register()
	return s
}

// Run server
func (s *Server) Run() error {
	return s.s.Run(s.bindAddr)
}

func (s *Server) register() {
	pprof.Register(s.s, "/debug")
}
