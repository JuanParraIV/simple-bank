package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/juanparraiv/simple-bank/db/sqlc"
	docs "github.com/juanparraiv/simple-bank/docs"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Server servers HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) *Server {

	server := &Server{store: store}
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/accounts")
		{
			eg.POST("/", server.createAccount)
			eg.GET("/", server.listAccounts)
			eg.GET(":id", server.getAccount)
		}
	}
	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// add routes to the server
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
// address: ":8080" for localhost:8080
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
