package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createEstadoReservaRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
}

type updateEstadoReservaRequest struct {
	IdEstadoReserva int32  `json:"idEstadoReserva" binding:"required"`
	Nombre          string `json:"nombre" binding:"required"`
	Descripcion     string `json:"descripcion"`
}

func (server *Server) createEstadoReserva(ctx *gin.Context) {
	var req createEstadoReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateEstadoReservaParams{
		Nombre: req.Nombre,
		Descripcion: sql.NullString{
			String: req.Descripcion,
			Valid:  req.Descripcion != "",
		},
	}

	result, err := server.dbtx.CreateEstadoReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{
		"idEstadoReserva": lastId,
		"message":         "EstadoReserva creado",
	})
}

func (server *Server) getAllEstadoReserva(ctx *gin.Context) {
	estados, err := server.dbtx.GetAllEstadoReserva(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, estados)
}

func (server *Server) getEstadoReservaById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	estado, err := server.dbtx.GetEstadoReservaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Estado de reserva no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, estado)
}

func (server *Server) updateEstadoReserva(ctx *gin.Context) {
	var req updateEstadoReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateEstadoReservaParams{
		Nombre: req.Nombre,
		Descripcion: sql.NullString{
			String: req.Descripcion,
			Valid:  req.Descripcion != "",
		},
		Idestadoreserva: req.IdEstadoReserva,
	}

	_, err := server.dbtx.UpdateEstadoReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Estado de reserva actualizado"})
}

func (server *Server) deleteEstadoReserva(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteEstadoReserva(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Estado de reserva eliminado"})
}