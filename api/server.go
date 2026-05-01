package api

import (
	"rest/dto"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type Server struct {
	dbtx   *dto.Queries
	router *gin.Engine
}

func NewServer(dbtx *dto.Queries) (*Server, error) {
	server := &Server{
		dbtx: dbtx,
	}
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,POST,PUT,DELETE",
		RequestHeaders:  "Origin,Authorization,Content_Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	//RUTAS SIN MIDDLEWARE
	router.POST("api/v1/chofer", server.createChofer)
	router.GET("api/v1/chofer", server.getAllChofer)
	router.GET("api/v1/chofer/:id", server.getChoferById)
	router.PUT("api/v1/chofer", server.updateChofer)
	router.DELETE("api/v1/chofer/:id", server.deleteChofer)
	router.POST("api/v1/ubicacion", server.createUbicacion)
	router.GET("api/v1/ubicacion", server.getAllUbicacion)
	router.GET("api/v1/ubicacion/:id", server.getUbicacionById)
	router.PUT("api/v1/ubicacion", server.updateUbicacion)
	router.DELETE("api/v1/ubicacion/:id", server.deleteUbicacion)

	//RUTAS CON MIDDLEWARE

	server.router = router
	return server, nil
}
func (server *Server) Start(url string) error {
	return server.router.Run(url)
}
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
