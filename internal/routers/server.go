package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ian995/UniqueBank/internal/repo"
)

type Server struct {
	store  repo.Store
	Router *gin.Engine
}

func NewServer(store repo.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id_account", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.Router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
