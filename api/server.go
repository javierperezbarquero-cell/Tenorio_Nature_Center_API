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
		RequestHeaders:  "Origin,Authorization,Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	//RUTAS SIN MIDDLEWARE
	//Login y registro
	router.POST("api/v1/usuario/login", server.login)
	router.POST("api/v1/usuario", server.createUsuario)
	//Manejo de imágenes de usuario
	router.POST("api/v1/usuario/upload", server.uploadUserImg)
	router.GET("api/v1/usuario/download/:filename", server.downloadUserImg)

	router.GET("api/v1/chofer", server.getAllChofer)
	router.GET("api/v1/chofer/:id", server.getChoferById)
	router.GET("api/v1/vehiculo", server.getAllVehiculos)
	router.GET("api/v1/vehiculo/:id", server.getVehiculoById)
	router.GET("api/v1/tour", server.getAllTours)
	router.GET("api/v1/tour/:id", server.getTourById)
	router.GET("api/v1/ubicacion", server.getAllUbicacion)
	router.GET("api/v1/ubicacion/:id", server.getUbicacionById)
	router.GET("api/v1/cliente", server.getAllClientes)
	router.GET("api/v1/cliente/:id", server.getClienteById)
	router.GET("api/v1/cliente/usuario/:id", server.getClienteByUsuario)
	router.GET("api/v1/guia", server.getAllGuia)
	router.GET("api/v1/guia/:id", server.getGuiaById)
	router.GET("api/v1/idioma", server.getAllIdioma)
	router.GET("api/v1/idioma/:id", server.getIdiomaById)
	router.GET("api/v1/emailcliente", server.getAllEmailCliente)
	router.GET("api/v1/emailcliente/:id", server.getEmailClienteById)
	router.GET("api/v1/idiomaguia", server.getAllIdiomaGuia)
	router.GET("api/v1/idiomaguia/idioma/:id", server.getIdiomaGuiaByIdioma)
	router.GET("api/v1/idiomaguia/guia/:id", server.getIdiomaGuiaByGuia)
	router.GET("api/v1/idiomaguia/:id", server.getIdiomaGuiaById)
	router.GET("api/v1/estadoreserva", server.getAllEstadoReserva)
	router.GET("api/v1/estadoreserva/:id", server.getEstadoReservaById)
	router.GET("api/v1/reserva", server.getAllReservas)
	router.GET("api/v1/reserva/:id", server.getReservaById)
	router.GET("api/v1/reserva/completa", server.createReservaCompleta)
	router.GET("api/v1/reserva/cliente/:idCliente")
	router.GET("api/v1/participante", server.getAllParticipantes)
	router.GET("api/v1/participante/:id", server.getParticipanteById)
	router.GET("api/v1/participante/cliente/:id", server.getParticipanteByCliente)
	router.GET("api/v1/participante/reserva/:id", server.getParticipanteByReserva)
	router.GET("api/v1/estadopago", server.getAllEstadoPago)
	router.GET("api/v1/estadopago/:id", server.getEstadoPagoById)
	router.GET("api/v1/factura", server.getAllFacturas)
	router.GET("api/v1/factura/:id", server.getFacturaById)
	router.GET("api/v1/facturaparticipante/factura/:id", server.getParticipantesByFactura)
	router.GET("api/v1/facturaparticipante/participante/:id", server.getFacturaByParticipante)
	router.GET("api/v1/empresacliente", server.getAllEmpresaCliente)
	router.GET("api/v1/empresacliente/:id", server.getEmpresaClienteById)
	router.GET("api/v1/detallereserva", server.getAllDetalleReserva)
	router.GET("api/v1/detallereserva/:id", server.getDetalleReservaById)

	//RUTAS CON MIDDLEWARE
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenBuilder))

	//Usuario
	authRoutes.GET("api/v1/usuario", server.getAllUsuarios)
	authRoutes.PUT("api/v1/usuario", server.updateUsuario)
	authRoutes.DELETE("api/v1/usuario/:id", server.deleteUsuario)
	//Chofer
	authRoutes.POST("api/v1/chofer", server.createChofer)
	authRoutes.PUT("api/v1/chofer", server.updateChofer)
	authRoutes.DELETE("api/v1/chofer/:id", server.deleteChofer)
	//Ubicacion
	authRoutes.POST("api/v1/ubicacion", server.createUbicacion)
	authRoutes.PUT("api/v1/ubicacion", server.updateUbicacion)
	authRoutes.DELETE("api/v1/ubicacion/:id", server.deleteUbicacion)
	//Vehiculo
	authRoutes.POST("api/v1/vehiculo", server.createVehiculo)
	authRoutes.PUT("api/v1/vehiculo", server.updateVehiculo)
	authRoutes.DELETE("api/v1/vehiculo/:id", server.deleteVehiculo)
	//Tour
	authRoutes.POST("api/v1/tour", server.createTour)
	authRoutes.PUT("api/v1/tour", server.updateTour)
	authRoutes.DELETE("api/v1/tour/:id", server.deleteTour)
	//Cliente
	authRoutes.POST("api/v1/cliente", server.createCliente)
	authRoutes.PUT("api/v1/cliente", server.updateCliente)
	authRoutes.DELETE("api/v1/cliente/:id", server.deleteCliente)
	//Guia
	authRoutes.POST("api/v1/guia", server.createGuia)
	authRoutes.PUT("api/v1/guia", server.updateGuia)
	authRoutes.DELETE("api/v1/guia/:id", server.deleteGuia)
	//Idioma
	authRoutes.POST("api/v1/idioma", server.createIdioma)
	authRoutes.PUT("api/v1/idioma", server.updateIdioma)
	authRoutes.DELETE("api/v1/idioma/:id", server.deleteIdioma)
	//EmailCliente
	authRoutes.POST("api/v1/emailcliente", server.createEmailCliente)
	authRoutes.PUT("api/v1/emailcliente", server.updateEmailCliente)
	authRoutes.DELETE("api/v1/emailcliente/:id", server.deleteEmailCliente)
	//IdiomaGuia
	authRoutes.POST("api/v1/idiomaguia", server.createIdiomaGuia)
	authRoutes.DELETE("api/v1/idiomaguia/:id", server.deleteIdiomaGuia)
	// Estado Reserva
	authRoutes.POST("api/v1/estadoreserva", server.createEstadoReserva)
	authRoutes.PUT("api/v1/estadoreserva", server.updateEstadoReserva)
	authRoutes.DELETE("api/v1/estadoreserva/:id", server.deleteEstadoReserva)
	// Reserva
	authRoutes.POST("api/v1/reserva", server.createReserva)
	authRoutes.PUT("api/v1/reserva", server.updateReserva)
	authRoutes.DELETE("api/v1/reserva/:id", server.deleteReserva)
	// Participante
	authRoutes.POST("api/v1/participante", server.createParticipante)
	authRoutes.DELETE("api/v1/participante/:id", server.deleteParticipante)
	// Estado Pago
	authRoutes.POST("api/v1/estadopago", server.createEstadoPago)
	authRoutes.PUT("api/v1/estadopago", server.updateEstadoPago)
	authRoutes.DELETE("api/v1/estadopago/:id", server.deleteEstadoPago)
	// Factura
	authRoutes.POST("api/v1/factura", server.createFactura)
	authRoutes.PUT("api/v1/factura", server.updateFactura)
	authRoutes.DELETE("api/v1/factura/:id", server.deleteFactura)
	// Factura Participante
	authRoutes.POST("api/v1/facturaparticipante", server.createFacturaParticipante)
	authRoutes.PUT("api/v1/facturaparticipante", server.updateFacturaParticipante)
	authRoutes.DELETE("api/v1/facturaparticipante/:id", server.deleteFacturaParticipante)
	// Empresa Cliente
	authRoutes.POST("api/v1/empresacliente", server.createEmpresaCliente)
	authRoutes.PUT("api/v1/empresacliente", server.updateEmpresaCliente)
	authRoutes.DELETE("api/v1/empresacliente/:id", server.deleteEmpresaCliente)
	// Detalle Reserva
	authRoutes.POST("api/v1/detallereserva", server.createDetalleReserva)
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
