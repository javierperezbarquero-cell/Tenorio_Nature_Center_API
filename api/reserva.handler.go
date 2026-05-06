package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createReservaRequest struct {
	IdCliente        int32  `json:"idCliente" binding:"required"`
	IdTour           int32  `json:"idTour" binding:"required"`
	IdGuia           int32  `json:"idGuia" binding:"required"`
	IdTransporte     int32  `json:"idTransporte" binding:"required"`
	IdUbicacion      int32  `json:"idUbicacion" binding:"required"`
	IdIdioma         int32  `json:"idIdioma" binding:"required"`
	IdEstadoReserva  int32  `json:"idEstadoReserva" binding:"required"`
	CantidadPersonas int32  `json:"cantidadPersonas" binding:"required"`
	FechaTour        string `json:"fechaTour" binding:"required"`
}

type updateReservaRequest struct {
	IdReserva        int32  `json:"idReserva" binding:"required"`
	IdCliente        int32  `json:"idCliente" binding:"required"`
	IdTour           int32  `json:"idTour" binding:"required"`
	IdGuia           int32  `json:"idGuia" binding:"required"`
	IdTransporte     int32  `json:"idTransporte" binding:"required"`
	IdUbicacion      int32  `json:"idUbicacion" binding:"required"`
	IdIdioma         int32  `json:"idIdioma" binding:"required"`
	IdEstadoReserva  int32  `json:"idEstadoReserva" binding:"required"`
	CantidadPersonas int32  `json:"cantidadPersonas" binding:"required"`
	FechaTour        string `json:"fechaTour" binding:"required"`
}

func (server *Server) createReserva(ctx *gin.Context) {
	var req createReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fechaTour, err := parsearFecha(req.FechaTour, "fechaTour")
	if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}

	args := dto.CreateReservaParams{
		Idcliente:        req.IdCliente,
		Idtour:           req.IdTour,
		Idguia:           req.IdGuia,
		Idtransporte:     req.IdTransporte,
		Idubicacion:      req.IdUbicacion,
		Ididioma:         req.IdIdioma,
		Idestadoreserva:  req.IdEstadoReserva,
		Cantidadpersonas: req.CantidadPersonas,
		Fechatour:        fechaTour,
	}

	result, err := server.dbtx.CreateReserva(ctx, args)
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
	fechaTour, err := parsearFecha(req.FechaTour, "fechaTour")
	if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}

	args := dto.UpdateReservaParams{
		Idcliente:        req.IdCliente,
		Idtour:           req.IdTour,
		Idguia:           req.IdGuia,
		Idtransporte:     req.IdTransporte,
		Idubicacion:      req.IdUbicacion,
		Ididioma:         req.IdIdioma,
		Idestadoreserva:  req.IdEstadoReserva,
		Cantidadpersonas: req.CantidadPersonas,
		Idreserva:        req.IdReserva,
		Fechatour:        fechaTour,
	}

	_, err = server.dbtx.UpdateReserva(ctx, args)
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
