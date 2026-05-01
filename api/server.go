package api

import (
	"rest/dto"
	"rest/security"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type Server struct {
	dbtx         *dto.Queries
	router       *gin.Engine
	tokenBuilder security.Builder
}

func NewServer(dbtx *dto.Queries, secret string) (*Server, error) {
	builder, err := security.NewPasetoBuilder(secret)
	if err != nil {
		return nil, err
	}

	server := &Server{
		dbtx:         dbtx,
		tokenBuilder: builder,
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
	router.POST("api/v1/login", server.login)
	router.POST("api/v1/usuario", server.createUsuario)
	router.PUT("api/v1/usuario", server.updateUsuario)
	router.DELETE("api/v1/usuario/:id", server.deleteUsuario)

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
