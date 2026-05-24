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
	//Login
	router.POST("api/v1/usuario/login", server.login)

	//RUTAS CON MIDDLEWARE
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenBuilder))

	//Usuario
	authRoutes.POST("api/v1/usuario", server.createUsuario)
	authRoutes.PUT("api/v1/usuario", server.updateUsuario)
	authRoutes.DELETE("api/v1/usuario/:id", server.deleteUsuario)
	//Chofer
	authRoutes.POST("api/v1/chofer", server.createChofer)
	authRoutes.GET("api/v1/chofer", server.getAllChofer)
	authRoutes.GET("api/v1/chofer/:id", server.getChoferById)
	authRoutes.PUT("api/v1/chofer", server.updateChofer)
	authRoutes.DELETE("api/v1/chofer/:id", server.deleteChofer)
	//Ubicacion
	authRoutes.POST("api/v1/ubicacion", server.createUbicacion)
	authRoutes.GET("api/v1/ubicacion", server.getAllUbicacion)
	authRoutes.GET("api/v1/ubicacion/:id", server.getUbicacionById)
	authRoutes.PUT("api/v1/ubicacion", server.updateUbicacion)
	authRoutes.DELETE("api/v1/ubicacion/:id", server.deleteUbicacion)
	//Vehiculo
	authRoutes.POST("api/v1/vehiculo", server.createVehiculo)
	authRoutes.GET("api/v1/vehiculo", server.getAllVehiculos)
	authRoutes.GET("api/v1/vehiculo/:id", server.getVehiculoById)
	authRoutes.PUT("api/v1/vehiculo", server.updateVehiculo)
	authRoutes.DELETE("api/v1/vehiculo/:id", server.deleteVehiculo)
	//Tour
	authRoutes.POST("api/v1/tour", server.createTour)
	authRoutes.GET("api/v1/tour", server.getAllTours)
	authRoutes.GET("api/v1/tour/:id", server.getTourById)
	authRoutes.PUT("api/v1/tour", server.updateTour)
	authRoutes.DELETE("api/v1/tour/:id", server.deleteTour)
	//Cliente
	authRoutes.POST("api/v1/cliente", server.createCliente)
	authRoutes.GET("api/v1/cliente", server.getAllClientes)
	authRoutes.GET("api/v1/cliente/:id", server.getClienteById)
	authRoutes.PUT("api/v1/cliente", server.updateCliente)
	authRoutes.DELETE("api/v1/cliente/:id", server.deleteCliente)
	//Guia
	authRoutes.POST("api/v1/guia", server.createGuia)
	authRoutes.GET("api/v1/guia", server.getAllGuia)
	authRoutes.GET("api/v1/guia/:id", server.getGuiaById)
	authRoutes.PUT("api/v1/guia", server.updateGuia)
	authRoutes.DELETE("api/v1/guia/:id", server.deleteGuia)
	//Idioma
	authRoutes.POST("api/v1/idioma", server.createIdioma)
	authRoutes.GET("api/v1/idioma", server.getAllIdioma)
	authRoutes.GET("api/v1/idioma/:id", server.getIdiomaById)
	authRoutes.PUT("api/v1/idioma", server.updateIdioma)
	authRoutes.DELETE("api/v1/idioma/:id", server.deleteIdioma)
	//EmailCliente
	authRoutes.POST("api/v1/emailcliente", server.createEmailCliente)
	authRoutes.GET("api/v1/emailcliente", server.getAllEmailCliente)
	authRoutes.GET("api/v1/emailcliente/:id", server.getEmailClienteById)
	authRoutes.PUT("api/v1/emailcliente", server.updateEmailCliente)
	authRoutes.DELETE("api/v1/emailcliente/:id", server.deleteEmailCliente)
	//IdiomaGuia
	authRoutes.POST("api/v1/idiomaguia", server.createIdiomaGuia)
	authRoutes.GET("api/v1/idiomaguia", server.getAllIdiomaGuia)
	authRoutes.GET("api/v1/idiomaguia/:id", server.getIdiomaGuiaById)
	authRoutes.DELETE("api/v1/idiomaguia/:id", server.deleteIdiomaGuia)
	// Estado Reserva
	authRoutes.POST("api/v1/estadoreserva", server.createEstadoReserva)
	authRoutes.GET("api/v1/estadoreserva", server.getAllEstadoReserva)
	authRoutes.GET("api/v1/estadoreserva/:id", server.getEstadoReservaById)
	authRoutes.PUT("api/v1/estadoreserva", server.updateEstadoReserva)
	authRoutes.DELETE("api/v1/estadoreserva/:id", server.deleteEstadoReserva)
	// Reserva
	authRoutes.POST("api/v1/reserva", server.createReserva)
	authRoutes.GET("api/v1/reserva", server.getAllReservas)
	authRoutes.GET("api/v1/reserva/:id", server.getReservaById)
	authRoutes.PUT("api/v1/reserva", server.updateReserva)
	authRoutes.DELETE("api/v1/reserva/:id", server.deleteReserva)
	// Participante
	authRoutes.POST("api/v1/participante", server.createParticipante)
	authRoutes.GET("api/v1/participante", server.getAllParticipantes)
	authRoutes.GET("api/v1/participante/:id", server.getParticipanteById)
	authRoutes.DELETE("api/v1/participante/:id", server.deleteParticipante)
	// Estado Pago
	authRoutes.POST("api/v1/estadopago", server.createEstadoPago)
	authRoutes.GET("api/v1/estadopago", server.getAllEstadoPago)
	authRoutes.GET("api/v1/estadopago/:id", server.getEstadoPagoById)
	authRoutes.PUT("api/v1/estadopago", server.updateEstadoPago)
	authRoutes.DELETE("api/v1/estadopago/:id", server.deleteEstadoPago)
	// Factura
	authRoutes.POST("api/v1/factura", server.createFactura)
	authRoutes.GET("api/v1/factura", server.getAllFacturas)
	authRoutes.GET("api/v1/factura/:id", server.getFacturaById)
	authRoutes.PUT("api/v1/factura", server.updateFactura)
	authRoutes.DELETE("api/v1/factura/:id", server.deleteFactura)
	// Empresa Cliente
	authRoutes.POST("api/v1/empresacliente", server.createEmpresaCliente)
	authRoutes.GET("api/v1/empresacliente", server.getAllEmpresaCliente)
	authRoutes.GET("api/v1/empresacliente/:id", server.getEmpresaClienteById)
	authRoutes.PUT("api/v1/empresacliente", server.updateEmpresaCliente)
	authRoutes.DELETE("api/v1/empresacliente/:id", server.deleteEmpresaCliente)
	// Detalle Reserva
	authRoutes.POST("api/v1/detallereserva", server.createDetalleReserva)
	authRoutes.GET("api/v1/detallereserva", server.getAllDetalleReserva)
	authRoutes.GET("api/v1/detallereserva/:id", server.getDetalleReservaById)
	authRoutes.PUT("api/v1/detallereserva", server.updateDetalleReserva)
	authRoutes.DELETE("api/v1/detallereserva/:id", server.deleteDetalleReserva)

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
