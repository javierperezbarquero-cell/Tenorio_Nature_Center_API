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
	//Chofer
	router.POST("api/v1/chofer", server.createChofer)
	router.GET("api/v1/chofer", server.getAllChofer)
	router.GET("api/v1/chofer/:id", server.getChoferById)
	router.PUT("api/v1/chofer", server.updateChofer)
	router.DELETE("api/v1/chofer/:id", server.deleteChofer)
	//Ubicacion
	router.POST("api/v1/ubicacion", server.createUbicacion)
	router.GET("api/v1/ubicacion", server.getAllUbicacion)
	router.GET("api/v1/ubicacion/:id", server.getUbicacionById)
	router.PUT("api/v1/ubicacion", server.updateUbicacion)
	router.DELETE("api/v1/ubicacion/:id", server.deleteUbicacion)
	//Vehiculo
	router.POST("api/v1/vehiculo", server.createVehiculo)
	router.GET("api/v1/vehiculo", server.getAllVehiculos)
	router.GET("api/v1/vehiculo/:id", server.getVehiculoById)
	router.PUT("api/v1/vehiculo", server.updateVehiculo)
	router.DELETE("api/v1/vehiculo/:id", server.deleteVehiculo)
	//Tour
	router.POST("api/v1/tour", server.createTour)
	router.GET("api/v1/tour", server.getAllTours)
	router.GET("api/v1/tour/:id", server.getTourById)
	router.PUT("api/v1/tour", server.updateTour)
	router.DELETE("api/v1/tour/:id", server.deleteTour)
	//Usuario
	router.POST("api/v1/usuario/login", server.login)
	router.POST("api/v1/usuario", server.createUsuario)
	router.PUT("api/v1/usuario", server.updateUsuario)
	router.DELETE("api/v1/usuario/:id", server.deleteUsuario)
	//Cliente
	router.POST("api/v1/cliente", server.createCliente)
	router.GET("api/v1/cliente", server.getAllClientes)
	router.GET("api/v1/cliente/:id", server.getClienteById)
	router.PUT("api/v1/cliente", server.updateCliente)
	router.DELETE("api/v1/cliente/:id", server.deleteCliente)
	//Guia
	router.POST("api/v1/guia", server.createGuia)
	router.GET("api/v1/guia", server.getAllGuia)
	router.GET("api/v1/guia/:id", server.getGuiaById)
	router.PUT("api/v1/guia", server.updateGuia)
	router.DELETE("api/v1/guia/:id", server.deleteGuia)
	//Idioma
	router.POST("api/v1/idioma", server.createIdioma)
	router.GET("api/v1/idioma", server.getAllIdioma)
	router.GET("api/v1/idioma/:id", server.getIdiomaById)
	router.PUT("api/v1/idioma", server.updateIdioma)
	router.DELETE("api/v1/idioma/:id", server.deleteIdioma)
	//EmailCliente
     router.POST("api/v1/emailcliente", server.createEmailCliente)
     router.GET("api/v1/emailcliente", server.getAllEmailCliente)
     router.GET("api/v1/emailcliente/:id", server.getEmailClienteById)
     router.PUT("api/v1/emailcliente", server.updateEmailCliente)
     router.DELETE("api/v1/emailcliente/:id", server.deleteEmailCliente)
     //IdiomaGuia
     router.POST("api/v1/idiomaguia", server.createIdiomaGuia)
     router.GET("api/v1/idiomaguia", server.getAllIdiomaGuia)
     router.GET("api/v1/idiomaguia/:id", server.getIdiomaGuiaById)
     router.PUT("api/v1/idiomaguia", server.updateIdiomaGuia)
     router.DELETE("api/v1/idiomaguia/:id", server.deleteIdiomaGuia)
     //Transporte
     router.POST("api/v1/transporte", server.createTransporte)
     router.GET("api/v1/transporte", server.getAllTransporte)
     router.GET("api/v1/transporte/:id", server.getTransporteById)
     router.PUT("api/v1/transporte", server.updateTransporte)
     router.DELETE("api/v1/transporte/:id", server.deleteTransporte)

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
