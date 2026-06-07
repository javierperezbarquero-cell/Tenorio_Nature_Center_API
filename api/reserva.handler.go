package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createReservaRequest struct {
	IdEstadoReserva int32 `json:"idEstadoReserva" binding:"required"`
}

type updateReservaRequest struct {
	IdReserva       int32 `json:"idReserva"       binding:"required"`
	IdEstadoReserva int32 `json:"idEstadoReserva" binding:"required"`
}

type participanteReservaRequest struct {
	Nombre        string `json:"nombre"`
	Identificador string `json:"identificador"`
	FechaNac      string `json:"fechaNac"`
	Telefono      string `json:"telefono"`
	Nacionalidad  string `json:"nacionalidad"`
}

type createReservaCompletaRequest struct {

	// cliente principal
	IdUsuario        *int32 `json:"idUsuario"`
	IdEmpresaCliente *int32 `json:"idEmpresaCliente"`

	NombreCliente        string `json:"nombreCliente"`
	IdentificadorCliente string `json:"identificadorCliente"`
	FechaNacCliente      string `json:"fechaNacCliente"`
	TelefonoCliente      string `json:"telefonoCliente"`
	NacionalidadCliente  string `json:"nacionalidadCliente"`

	// reserva
	IdEstadoReserva int32 `json:"idEstadoReserva"`

	// detalle reserva
	IdTour      int32  `json:"idTour"`
	IdGuia      int32  `json:"idGuia"`
	IdChofer    int32  `json:"idChofer"`
	IdUbicacion int32  `json:"idUbicacion"`
	IdIdioma    int32  `json:"idIdioma"`
	FechaTour   string `json:"fechaTour"`
	Precio      string `json:"precio"`

	// acompañantes
	Participantes []participanteReservaRequest `json:"participantes"`
}

func (server *Server) createReserva(ctx *gin.Context) {
	var req createReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.dbtx.CreateReserva(ctx, req.IdEstadoReserva)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idReserva": lastId,
		"message":   "Reserva creada",
	})
}

func (server *Server) getAllReservas(ctx *gin.Context) {
	reservas, err := server.dbtx.GetAllReserva(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, reservas)
}

func (server *Server) getReservaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	reserva, err := server.dbtx.GetReservaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Reserva no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, reserva)
}

func (server *Server) updateReserva(ctx *gin.Context) {
	var req updateReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateReservaParams{
		Idestadoreserva: req.IdEstadoReserva,
		Idreserva:       req.IdReserva,
	}

	_, err := server.dbtx.UpdateReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva actualizada"})
}

func (server *Server) deleteReserva(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteReserva(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva eliminada"})
}

func (server *Server) createReservaCompleta(ctx *gin.Context) {

	var req createReservaCompletaRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fechaTour, err := parsearFecha(req.FechaTour, "fechaTour")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fechaRegistro := fechaTour

	fechaNacCliente, err := parsearFecha(
		req.FechaNacCliente,
		"fechaNacCliente",
	)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//---------------------------------------------------
	// CLIENTE PRINCIPAL
	//---------------------------------------------------

	clienteResult, err := server.dbtx.CreateCliente(
		ctx,
		dto.CreateClienteParams{
			Idempresacliente: toNullInt32(req.IdEmpresaCliente),
			Idusuario:        toNullInt32(req.IdUsuario),

			Nombre:        req.NombreCliente,
			Identificador: req.IdentificadorCliente,
			Fechanac:      fechaNacCliente,
			Telefono:      req.TelefonoCliente,
			Nacionalidad:  req.NacionalidadCliente,
			Fecharegistro: fechaRegistro,
		},
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	idCliente64, _ := clienteResult.LastInsertId()
	idCliente := int32(idCliente64)

	//---------------------------------------------------
	// RESERVA
	//---------------------------------------------------

	reservaResult, err := server.dbtx.CreateReserva(
		ctx,
		req.IdEstadoReserva,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	idReserva64, _ := reservaResult.LastInsertId()
	idReserva := int32(idReserva64)

	//---------------------------------------------------
	// DETALLE RESERVA
	//---------------------------------------------------

	_, err = server.dbtx.CreateDetalleReserva(
		ctx,
		dto.CreateDetalleReservaParams{
			Idreserva:      idReserva,
			Idtour:         req.IdTour,
			Idguia:         req.IdGuia,
			Idchofer:       req.IdChofer,
			Idubicacion:    req.IdUbicacion,
			Ididioma:       req.IdIdioma,
			Fechatour:      fechaTour,
			Preciounitario: req.Precio,
		},
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//---------------------------------------------------
	// PARTICIPANTE PRINCIPAL
	//---------------------------------------------------

	_, err = server.dbtx.CreateParticipante(
		ctx,
		dto.CreateParticipanteParams{
			Idcliente: idCliente,
			Idreserva: idReserva,
		},
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//---------------------------------------------------
	// PARTICIPANTES EXTRA
	//---------------------------------------------------

	for _, p := range req.Participantes {

		fechaNac, err := parsearFecha(
			p.FechaNac,
			"fechaNac",
		)

		if err != nil {
			continue
		}

		resultCliente, err := server.dbtx.CreateCliente(
			ctx,
			dto.CreateClienteParams{
				Nombre:           p.Nombre,
				Identificador:    p.Identificador,
				Fechanac:         fechaNac,
				Telefono:         p.Telefono,
				Nacionalidad:     p.Nacionalidad,
				Fecharegistro:    fechaRegistro,
				Idusuario:        sql.NullInt32{},
				Idempresacliente: sql.NullInt32{},
			},
		)

		if err != nil {
			continue
		}

		idExtra64, _ := resultCliente.LastInsertId()

		server.dbtx.CreateParticipante(
			ctx,
			dto.CreateParticipanteParams{
				Idcliente: int32(idExtra64),
				Idreserva: idReserva,
			},
		)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"idReserva": idReserva,
		"message":   "Reserva creada correctamente",
	})
}
