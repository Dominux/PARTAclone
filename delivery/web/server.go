package web

import (
	"github.com/gin-gonic/gin"
)


type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{ engine: gin.Default() }
}

func (s *Server) Run(port string) {
	s.engine.Run(":" + port)
}

//////////////////////////////////////////////
////	HTTP Methods	
//////////////////////////////////////////////

func (s *Server) GET(path string, handlers ...gin.HandlerFunc) {
	s.engine.GET(path, handlers...)
}

func (s *Server) POST(path string, handlers ...gin.HandlerFunc) {
	s.engine.POST(path, handlers...)
}

func (s *Server) PUT(path string, handlers ...gin.HandlerFunc) {
	s.engine.PUT(path, handlers...)
}

func (s *Server) DELETE(path string, handlers ...gin.HandlerFunc) {
	s.engine.DELETE(path, handlers...)
}

func (s *Server) PATCH(path string, handlers ...gin.HandlerFunc) {
	s.engine.PATCH(path, handlers...)
}

func (s *Server) HEAD(path string, handlers ...gin.HandlerFunc) {
	s.engine.HEAD(path, handlers...)
}

func (s *Server) OPTIONS(path string, handlers ...gin.HandlerFunc) {
	s.engine.OPTIONS(path, handlers...)
}
