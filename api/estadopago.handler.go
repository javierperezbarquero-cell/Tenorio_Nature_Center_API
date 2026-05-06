package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createEstadoPagoRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
}

type updateEstadoPagoRequest struct {
	IdEstadoPago int32  `json:"idEstadoPago" binding:"required"`
	Nombre       string `json:"nombre" binding:"required"`
	Descripcion  string `json:"descripcion"`
}

func (server *Server) createEstadoPago(ctx *gin.Context) {
	var req createEstadoPagoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateEstadoPagoParams{
		Nombre: req.Nombre,
		Descripcion: sql.NullString{
			String: req.Descripcion,
			Valid:  req.Descripcion != "",
		},
	}

	result, err := server.dbtx.CreateEstadoPago(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idEstadoPago": lastId,
		"message":      "EstadoPago creado",
	})
}

func (server *Server) getAllEstadoPago(ctx *gin.Context) {
	estados, err := server.dbtx.GetAllEstadoPago(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, estados)
}

func (server *Server) getEstadoPagoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	estado, err := server.dbtx.GetEstadoPagoById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Estado de pago no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, estado)
}

func (server *Server) updateEstadoPago(ctx *gin.Context) {
	var req updateEstadoPagoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateEstadoPagoParams{
		Nombre: req.Nombre,
		Descripcion: sql.NullString{
			String: req.Descripcion,
			Valid:  req.Descripcion != "",
		},
		Idestadopago: req.IdEstadoPago,
	}

	_, err := server.dbtx.UpdateEstadoPago(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Estado de pago actualizado"})
}

func (server *Server) deleteEstadoPago(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteEstadoPago(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Estado de pago eliminado"})
}