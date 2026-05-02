package api

import (
	"database/sql"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createTransporteRequest struct {
	IdChofer int32 `json:"idChofer" binding:"required"`
}

type updateTransporteRequest struct {
	IdTransporte int32  `json:"idTransporte"   binding:"required"`
	IdChofer     int32  `json:"idChofer"      binding:"required"`
}

func (server *Server) createTransporte(ctx *gin.Context) {
	var req createTransporteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	transporte, err := server.dbtx.CreateTransporte(ctx, req.IdChofer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	lastId, _ := transporte.LastInsertId()
	ctx.JSON(http.StatusOK, gin.H{"generated_id": lastId})
}

func (server *Server) getAllTransporte(ctx *gin.Context) {
	transporte, err := server.dbtx.GetAllTransporte(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transporte)
}

func (server *Server) getTransporteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	transporte, err := server.dbtx.GetTransporteById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Transporte no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transporte)
}

func (server *Server) updateTransporte(ctx *gin.Context) {
	var req updateTransporteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateTransporteParams{
		Idchofer:     req.IdChofer,
		Idtransporte: req.IdTransporte,
	}

	_, err := server.dbtx.UpdateTransporte(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Transporte actualizado correctamente"})
}

func (server *Server) deleteTransporte(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = server.dbtx.DeleteTransporte(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Transporte eliminado correctamente"})
}
