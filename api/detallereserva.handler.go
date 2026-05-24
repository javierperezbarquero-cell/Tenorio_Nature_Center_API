package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createDetalleReservaRequest struct {
	IdReserva     int32  `json:"idReserva"     binding:"required"`
	IdTour        int32  `json:"idTour"        binding:"required"`
	IdGuia        int32  `json:"idGuia"        binding:"required"`
	IdChofer      int32  `json:"idChofer"      binding:"required"`
	IdUbicacion   int32  `json:"idUbicacion"   binding:"required"`
	IdIdioma      int32  `json:"idIdioma"      binding:"required"`
	FechaTour     string `json:"fechaTour"     binding:"required"`
	PrecioUnitario string `json:"precioUnitario" binding:"required"`
}

type updateDetalleReservaRequest struct {
	IdDetalleReserva int32  `json:"idDetalleReserva" binding:"required"`
	IdTour           int32  `json:"idTour"           binding:"required"`
	IdGuia           int32  `json:"idGuia"           binding:"required"`
	IdChofer         int32  `json:"idChofer"         binding:"required"`
	IdUbicacion      int32  `json:"idUbicacion"      binding:"required"`
	IdIdioma         int32  `json:"idIdioma"         binding:"required"`
	FechaTour        string `json:"fechaTour"        binding:"required"`
	PrecioUnitario   string `json:"precioUnitario"   binding:"required"`
}

func (server *Server) createDetalleReserva(ctx *gin.Context) {
	var req createDetalleReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fechaTour, err := parsearFecha(req.FechaTour, "fechaTour")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validarDecimal(req.PrecioUnitario, "precioUnitario"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.CreateDetalleReservaParams{
		Idreserva:      req.IdReserva,
		Idtour:         req.IdTour,
		Idguia:         req.IdGuia,
		Idchofer:       req.IdChofer,
		Idubicacion:    req.IdUbicacion,
		Ididioma:       req.IdIdioma,
		Fechatour:      fechaTour,
		Preciounitario: req.PrecioUnitario,
	}

	result, err := server.dbtx.CreateDetalleReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idDetalleReserva": lastId,
		"message":          "Detalle de reserva creado",
	})
}

func (server *Server) getAllDetalleReserva(ctx *gin.Context) {
	detalles, err := server.dbtx.GetAllDetalleReserva(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, detalles)
}

func (server *Server) getDetalleReservaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	detalle, err := server.dbtx.GetDetalleReservaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Detalle de reserva no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, detalle)
}

func (server *Server) getDetalleReservaByReserva(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	detalles, err := server.dbtx.GetDetalleReservaByReserva(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Detalles de la reserva no encontrados"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, detalles)
}

func (server *Server) updateDetalleReserva(ctx *gin.Context) {
	var req updateDetalleReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fechaTour, err := parsearFecha(req.FechaTour, "fechaTour")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validarDecimal(req.PrecioUnitario, "precioUnitario"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := dto.UpdateDetalleReservaParams{
		Idtour:           req.IdTour,
		Idguia:           req.IdGuia,
		Idchofer:         req.IdChofer,
		Idubicacion:      req.IdUbicacion,
		Ididioma:         req.IdIdioma,
		Fechatour:        fechaTour,
		Preciounitario:   req.PrecioUnitario,
		Iddetallereserva: req.IdDetalleReserva,
	}

	_, err = server.dbtx.UpdateDetalleReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Detalle de reserva actualizado"})
}

func (server *Server) deleteDetalleReserva(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteDetalleReserva(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Detalle de reserva eliminado"})
}